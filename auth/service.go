package auth

import (
	// 用于处理上下文

	"context"
	"encoding/hex" // 十六进制编码
	"errors"       // 错误处理
	"fmt"          // 用于管理ECDSA私钥
	"log"          // 日志
	"math/big"     // 处理大整数
	"net/http"     // HTTP处理
	"strconv"      // 字符串和数字转换
	"strings"      // 字符串处理
	"time"

	"golang.org/x/crypto/sha3" // Keccak256 哈希算法

	// 时间处理
	"github.com/ethereum/go-ethereum/accounts/abi/bind" // 绑定智能合约
	"github.com/ethereum/go-ethereum/common"            // 处理以太坊地址和哈希
	"github.com/ethereum/go-ethereum/crypto"            // 加载私钥和生成签名

	"github.com/ethereum/go-ethereum/ethclient" // 连接以太坊节点

	"project-root/transaction"

	"github.com/golang-jwt/jwt/v4"

	"github.com/gin-gonic/gin" // Gin框架用于API开发
	cad "github.com/seencxy/lsr/common"
	"golang.org/x/crypto/bcrypt" // 密码哈希处理
	"gorm.io/gorm"               // ORM数据库处理
	"gorm.io/gorm/logger"        // 数据库日志处理
)

type AuthService struct {
	db     *gorm.DB
	client *ethclient.Client // 区块链客户端实例
}

func NewAuthService(db *gorm.DB, ethClient *ethclient.Client) *AuthService {
	// 打开数据库 SQL 日志记录以便调试
	db.Logger = logger.Default.LogMode(logger.Info) // 打印详细的 SQL 日志
	return &AuthService{
		db:     db,
		client: ethClient,
	}
}

// Register 方法，用于用户注册
func (s *AuthService) Register(user *User) error {
	// 检查密码是否已经被哈希
	if !IsPasswordHashed(user.Password) { // 这里修正逻辑，只有密码未被哈希时才进行哈希
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(strings.TrimSpace(user.Password)), bcrypt.DefaultCost)
		if err != nil {
			log.Println("Error generating password hash:", err)
			return err
		}
		user.PasswordHash = string(hashedPassword)
	} else {
		log.Println("Password is already hashed, skipping hashing process")
		user.PasswordHash = user.Password
	}

	// 创建用户记录，保存到数据库
	if err := s.db.Create(user).Error; err != nil {
		return err
	}

	// 需要为用户分配一个以太坊地址
	_, addr, err := cad.CreateOneAddress()
	if err != nil {
		return err
	}
	newAccount := Account{
		UserID:          user.UserID,
		EthereumAddress: addr,
		Balance:         0,
		AccountType:     "Receiver",
		AccountStatus:   "Active",
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
	return s.db.Create(&newAccount).Error
}

// IsPasswordHashed 用于检查密码是否已经被哈希
func IsPasswordHashed(password string) bool {
	// bcrypt 哈希格式通常以 $2a$ 开头
	return len(password) > 0 && password[:4] == "$2a$"
}

// Login 方法，用于用户登录
func (s *AuthService) Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// 检查 JSON 输入
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// 查找数据库中的用户
	var user User
	if err := s.db.Where("Email = ?", input.Email).First(&user).Error; err != nil {
		log.Println("User not found:", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// 增加调试信息，打印用户输入的密码和数据库中的哈希密码
	log.Println("Input Email (Login):", input.Email)
	log.Println("Input Password (Login):", input.Password)
	log.Println("Stored Hashed Password (DB):", user.PasswordHash)

	// 使用 bcrypt 验证明文密码与数据库中的哈希密码
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(strings.TrimSpace(input.Password)))
	if err != nil {
		log.Println("Password mismatch:", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// 如果密码验证成功，生成 JWT
	token, err := s.GenerateToken(&user)
	if err != nil {
		log.Println("JWT generation failed:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	log.Println("Login successful, JWT generated")
	// 登录成功，返回 JWT 给用户
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
	})
}

// ValidateToken method to validate JWT
func (s *AuthService) ValidateToken(tokenString string) (map[string]interface{}, error) {
	claims, err := ParseJWT(tokenString)
	if err != nil {
		return nil, err
	}
	return claims, nil
}

func ParseJWT(tokenString string) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Add your signing key here
		return []byte("your-secret-key"), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, errors.New("invalid token")
	}
}

// SelectUserById 根据用户ID查询用户以太坊地址
func SelectUserById(userId int, db *gorm.DB) (string, error) {
	var account Account
	if tx := db.Where("UserID = ?", userId).First(&account); tx.Error != nil {
		return "", tx.Error
	}
	return account.EthereumAddress, nil
}

// GenerateToken 方法，生成 JWT
func (s *AuthService) GenerateToken(user *User) (string, error) {
	// 调用 jwt.go 中的 GenerateJWT 方法生成 JWT
	return GenerateJWT(user)
}

// 验证密码
func (s *AuthService) VerifyPassword(hashedPassword, oldPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(oldPassword))
}

// 更新用户密码
func (s *AuthService) UpdatePassword(userID uint, newPassword string) error {
	var user User

	if err := s.db.First(&user, userID).Error; err != nil {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.PasswordHash = string(hashedPassword)
	return s.db.Save(&user).Error
}

// CallSmartContractTransfer 负责与区块链交互并执行交易记录
func CurrencyTypeToUint8(currency string) (uint8, error) {
	switch currency {
	case "USD":
		return 0, nil
	case "EUR":
		return 1, nil
	case "ETH":
		return 2, nil
	default:
		return 0, fmt.Errorf("unsupported currency type: %s", currency)
	}
}

// 假设这里是银行账户与以太坊地址的映射关系
var bankToEthMap = map[string]string{
	"12345": "0xAb8483F64d9C6d1EcF9b849Ae677dD3315835Cb2",
	"23456": "0x4B0897b0513fdc7C541B6d9D7E929C4e5364D2dB",
}

// BankAccountToEthereumAddress 将银行账户映射为以太坊地址
func (s *AuthService) BankAccountToEthereumAddress(bankAccount string) (string, error) {
	fmt.Println("银行账户为")
	fmt.Println(bankAccount)
	ethAddress, exists := bankToEthMap[bankAccount]
	if !exists {
		return "", fmt.Errorf("银行账户未映射到以太坊地址: %s", bankAccount)
	}
	return ethAddress, nil
}

// CallSmartContractTransfer 调用智能合约，执行交易
func (s *AuthService) CallSmartContractTransfer(req *TransactionRequest) (string, error) {
	// 初始化智能合约
	contract, err := s.initContract()
	if err != nil {
		return "", err
	}

	// 解析私钥
	privateKeyHex := "20384cab0f60a10ecec491f2fc9e2715bd78856e7cfee788a1a1cd0dc641d133"
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return "", fmt.Errorf("解析私钥失败: %v", err)
	}

	//Error recording transaction on-chain: unsupported currency type:      不支持的币种，，
	// 创建交易选项
	chainID := big.NewInt(1337) // Ganache 默认链 ID
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return "", fmt.Errorf("创建交易选项失败: %v", err)
	}

	// 设置交易选项参数（如Gas限制、Gas价格）
	auth.GasLimit = uint64(300000)          // 设置 Gas 上限
	auth.GasPrice = big.NewInt(20000000000) // 设置 Gas 价格

	// 转换交易金额为 Wei
	amountInWei := new(big.Int)
	amountInWei.SetString(fmt.Sprintf("%.0f", req.Amount*1e18), 10)

	// 将货币类型转换为 uint8
	currencyUint8, err := CurrencyTypeToUint8(req.CurrencyType)
	if err != nil {
		fmt.Println("货币类型为")
		fmt.Println(req.CurrencyType)
		return "", fmt.Errorf("货币 转义%s", err)
	}
	// 获取以太坊地址
	fmt.Println("获取")
	fmt.Println(req.SenderAccount)
	// senderEthAddress, err := s.BankAccountToEthereumAddress(req.SenderAccount)
	// if err != nil {
	// 	return "", fmt.Errorf("无效的发送者账户地址: %v", err)
	// }

	// receiverEthAddress, err := s.BankAccountToEthereumAddress(req.ReceiverAccount)
	// if err != nil {
	// 	return "", fmt.Errorf("无效的接收者账户地址: %v", err)
	// }

	// 生成唯一的 transactionId
	transactionId := fmt.Sprintf("txn-%d", time.Now().UnixNano())

	// 执行智能合约交易
	tx, err := contract.RecordTransaction(
		auth,
		transactionId,
		req.SenderAccount,
		req.ReceiverAccount,
		amountInWei,
		currencyUint8,
		big.NewInt(time.Now().Unix()),
		req.Remarks,
	)

	if err != nil {
		return "", fmt.Errorf("智能合约交易失败: %v", err)
	}
	// 检查交易是否成功提交
	receipt, err := bind.WaitMined(context.Background(), s.client, tx)
	if err != nil {
		return "", fmt.Errorf("等待交易确认失败: %v", err)
	}
	if receipt.Status == 0 {
		return "", fmt.Errorf("交易失败，未成功提交到区块链")
	}

	// 返回交易哈希
	return tx.Hash().Hex(), nil
}

// initContract 初始化智能合约
func (s *AuthService) initContract() (*transaction.FinancialTransaction, error) {
	// 连接到区块链客户端
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		return nil, fmt.Errorf("连接区块链失败: %v", err)
	}

	// 将客户端保存到 AuthService 结构体中
	s.client = client

	// 初始化智能合约实例
	contractAddress := common.HexToAddress("0x6b2D7136543a8850f9237a6460C535882f0d3f57")
	return transaction.NewFinancialTransaction(contractAddress, client)
}

// Authenticate 方法，根据 email 和 password 验证用户
func (s *AuthService) Authenticate(email, password string) (*User, error) {
	var user User
	if err := s.db.Where("Email = ?", email).First(&user).Error; err != nil {
		return nil, err // 如果用户不存在，返回错误
	}

	// 检查密码是否匹配
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return nil, errors.New("invalid password") // 密码不正确
	}

	return &user, nil // 返回用户
}

// 获取账户余额
func (s *AuthService) GetBalance(address string) (float64, error) {
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		return 0, err
	}

	balance, err := client.BalanceAt(context.Background(), common.HexToAddress(address), nil)
	if err != nil {
		return 0, err
	}

	balanceInEth := new(big.Float).Quo(new(big.Float).SetInt(balance), big.NewFloat(1e18))
	result, _ := balanceInEth.Float64()
	return result, nil
}

// 获取指定地址的最近交易记录
func (s *AuthService) GetRecentTransactionsByAddress(uid int) ([]TransactionRecord, error) {
	var transactions []TransactionRecord
	if err := s.db.Where("UserID = ? ", uid).
		Order("timestamp desc").Limit(10).Find(&transactions).Error; err != nil {
		return nil, fmt.Errorf("查询交易记录失败: %v", err)
	}
	return transactions, nil
}

// generateAddressFromAccount 将银行账户 int 转换为以太坊地址格式
func generateAddressFromAccount(account int) (string, error) {
	// 将 int 转换为 string
	accountStr := strconv.Itoa(account)

	// 使用 Keccak256 哈希算法生成以太坊地址
	hash := sha3.NewLegacyKeccak256()
	hash.Write([]byte(accountStr))
	hashed := hash.Sum(nil)

	// 以太坊地址为哈希结果的最后 20 个字节，转换为十六进制字符串
	address := "0x" + hex.EncodeToString(hashed[len(hashed)-20:])
	return address, nil
}

// RecordTransaction 调用智能合约并存储交易记录
func (s *AuthService) RecordTransaction(req *TransactionRequest) (string, error) {
	// 调用智能合约，并生成交易哈希
	fmt.Println("发送方账户为")
	fmt.Println(req.SenderAccount)
	txHash, err := s.CallSmartContractTransfer(req)
	if err != nil {
		return "", fmt.Errorf("智能合约调用失败: %v", err)
	}
	sa, _ := strconv.Atoi(req.SenderAccount)
	// 将 SenderAccount 和 ReceiverAccount 转换为以太坊地址
	senderAddress, err := generateAddressFromAccount(sa)
	if err != nil {
		return "", fmt.Errorf("生成 senderAddress 失败: %v", err)
	}
	receiverAddress, err := generateAddressFromAccount(sa)
	if err != nil {
		return "", fmt.Errorf("生成 receiverAddress 失败: %v", err)
	}

	// 创建 TransactionRecord 并存储到数据库
	txRecord := TransactionRecord{
		TxID:            txHash,
		SenderAccount:   senderAddress,   // 以太坊地址作为 string 存储
		ReceiverAccount: receiverAddress, // 以太坊地址作为 string 存储
		Amount:          req.Amount,
		CurrencyType:    req.CurrencyType,
		CurrencyUnit:    req.CurrencyUnit,
		TxHash:          txHash,
		BlockHeight:     0, // 根据需要设置实际的区块高度
		Note:            req.Remarks,
		Status:          "success",
		Timestamp:       time.Now(),
		CreatedAt:       time.Now(),
	}

	// 保存交易记录到数据库
	if err := s.db.Create(&txRecord).Error; err != nil {
		return "", fmt.Errorf("保存交易记录失败: %v", err)
	}

	return txHash, nil
}
