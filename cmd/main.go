package main

import (
	"fmt"
	"log"
	"math/big"
	"net/http"
	"project-root/auth"
	"project-root/transaction"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// 初始化数据库和区块链客户端
	db := initDB()
	defer closeDB(db)
	client := initEthereumClient()
	defer client.Close()

	// 初始化Gin路由器
	router := gin.Default()
	router.Use(auth.CORSMiddleware())
	router.Static("/frontend", "./frontend")

	// 初始化Auth服务和处理器
	authService := auth.NewAuthService(db, client)
	authHandler := auth.NewAuthHandler(authService)

	// 配置公共路由
	initPublicRoutes(router, authHandler)

	// 配置受保护路由
	protected := router.Group("/api", auth.JWTAuthMiddleware())
	initProtectedRoutes(protected, db, authHandler)

	// 启动服务器
	router.Run(":8080")
}

// 初始化数据库连接
func initDB() *gorm.DB {
	dsn := "root:20240410Mw@@tcp(127.0.0.1:3306)/crossborderpayment?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	db.AutoMigrate(&auth.TransactionRecord{})
	return db
}

func closeDB(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err == nil {
		sqlDB.Close()
	}
}

// 初始化Ethereum客户端和智能合约实例
func initEthereumClient() *ethclient.Client {
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatalf("Failed to connect to Ethereum client: %v", err)
	}
	return client
}

func loadPrivateKey() (*bind.TransactOpts, error) {
	privateKey, err := crypto.HexToECDSA("20384cab0f60a10ecec491f2fc9e2715bd78856e7cfee788a1a1cd0dc641d133")
	if err != nil {
		return nil, fmt.Errorf("Failed to parse private key: %v", err)
	}
	chainID := big.NewInt(1337)
	return bind.NewKeyedTransactorWithChainID(privateKey, chainID)
}

// 配置公共路由
func initPublicRoutes(router *gin.Engine, authHandler *auth.AuthHandler) {
	router.POST("/register", authHandler.Register)
	router.POST("/login", authHandler.Login)
}

// 配置受保护路由
func initProtectedRoutes(protected *gin.RouterGroup, db *gorm.DB, authHandler *auth.AuthHandler) {
	mainService := mainServer{
		db,
	}
	protected.GET("/transactions", authHandler.GetTransactions)
	protected.POST("/record-transaction", mainService.recordTransactionHandler)
	protected.GET("/transaction/:id", getTransactionHandler(db))
	protected.GET("/profile", authHandler.GetProfile)
	protected.POST("/change-password", authHandler.ChangePassword)
	protected.GET("/account-balance", mainService.accountBalanceHandler)
	protected.POST("/submit-transaction", authHandler.SubmitTransaction)
}

type mainServer struct {
	Db *gorm.DB
}

// 记录交易并保存到数据库的处理器
func (m *mainServer) recordTransactionHandler(c *gin.Context) {
	var req struct {
		TransactionID   string  `json:"transaction_id"`
		SenderAccount   string  `json:"sender_account"`
		ReceiverAccount string  `json:"receiver_account"`
		Amount          float64 `json:"amount"`
		Currency        string  `json:"currency"`
		Note            string  `json:"note"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	currencyEnum, err := mapCurrencyToEnum(req.Currency)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid currency"})
		return
	}

	senderAddress := generateAddressFromAccount(req.SenderAccount)
	receiverAddress := generateAddressFromAccount(req.ReceiverAccount)
	amountBigInt, _ := new(big.Float).Mul(big.NewFloat(req.Amount), big.NewFloat(1e18)).Int(nil)
	timestamp := big.NewInt(time.Now().Unix())

	auth1, err := loadPrivateKey()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load private key"})
		return
	}
	auth1.From = senderAddress

	contract, client := initContract()
	defer client.Close()
	tx, err := contract.RecordTransaction(auth1, req.TransactionID, req.SenderAccount, req.ReceiverAccount, amountBigInt, currencyEnum, timestamp, req.Note)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to record transaction"})
		return
	}

	record := auth.TransactionRecord{
		TxID:            req.TransactionID,
		SenderAccount:   senderAddress.Hex(),
		ReceiverAccount: receiverAddress.Hex(),
		Amount:          req.Amount,
		CurrencyType:    req.Currency,
		TxHash:          tx.Hash().Hex(),
		Timestamp:       time.Now(),
	}

	if err := m.Db.Create(&record).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save transaction"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"transactionHash": tx.Hash().Hex(),
		"transactionID":   req.TransactionID,
	})
}

func getTransactionHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		txID := c.Param("id")
		var record auth.TransactionRecord
		if err := db.Where("tx_hash = ?", txID).First(&record).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"transaction": record})
	}
}

func (m *mainServer) accountBalanceHandler(c *gin.Context) {
	accountAddress := "0xBEe92eBed029362F1CE813a9cBF2bC32Bd2c4324"
	client := initEthereumClient()
	defer client.Close()

	if !common.IsHexAddress(accountAddress) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid account address"})
		return
	}

	balance, err := client.BalanceAt(c, common.HexToAddress(accountAddress), nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get account balance"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"balance": balance.String()})
}

func submitTransactionHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request struct {
			Account string `json:"account"`
		}

		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		var transactions []auth.TransactionRecord
		if err := db.Where("SenderAccount = ? OR ReceiverAccount = ?", request.Account, request.Account).Find(&transactions).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch transactions"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"transactions": transactions})
	}
}

// 辅助函数
func mapCurrencyToEnum(currency string) (uint8, error) {
	switch currency {
	case "ETH":
		return 0, nil
	case "BTC":
		return 1, nil
	case "USD":
		return 2, nil
	default:
		return 0, fmt.Errorf("invalid currency: %s", currency)
	}
}

func generateAddressFromAccount(account string) common.Address {
	return common.HexToAddress(fmt.Sprintf("%x", account))
}

// initContract 初始化并返回智能合约实例和以太坊客户端
func initContract() (*transaction.FinancialTransaction, *ethclient.Client) {
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatal("Failed to connect to Ethereum client:", err)
	}

	contractAddress := "0x6b2D7136543a8850f9237a6460C535882f0d3f57"
	//contractAddress := "0xaa9709Cc41c2805523fF39Cb8C7B6e0d85C3036f"
	address := common.HexToAddress(contractAddress)

	contract, err := transaction.NewFinancialTransaction(address, client)
	if err != nil {
		log.Fatal("Failed to initialize contract:", err)
	}
	return contract, client
}
