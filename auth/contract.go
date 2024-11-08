package auth

import (
	"context"
	"errors"
	"log"
	"math/big"
	"os"                       // 从环境变量中读取私钥
	"project-root/transaction" // 确保路径与实际项目结构匹配
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// CallSmartContractTransfer 调用智能合约的 RecordTransaction 方法
func CallSmartContractTransfer(
	sender, receiver string, amount int64, note, txID string, currencyType uint8,
) (string, error) {

	// 连接到本地以太坊节点 (Ganache)
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Printf("连接以太坊客户端失败: %v", err)
		return "", errors.New("无法连接以太坊节点")
	}
	defer client.Close() // 确保关闭客户端连接

	// 加载合约地址并创建合约实例
	contractAddress := common.HexToAddress("0x6b2D7136543a8850f9237a6460C535882f0d3f57") // 替换为你的合约地址
	instance, err := transaction.NewFinancialTransaction(contractAddress, client)
	if err != nil {
		log.Printf("创建合约实例失败: %v", err)
		return "", errors.New("创建合约实例失败")
	}

	// 从环境变量中读取私钥 (确保私钥保密)
	privateKeyHex := os.Getenv("ETH_PRIVATE_KEY")
	if privateKeyHex == "" {
		return "", errors.New("未找到私钥，请检查环境变量设置")
	}
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Printf("解析私钥失败: %v", err)
		return "", errors.New("解析私钥失败")
	}

	// 创建交易选项并使用私钥签名
	chainID := big.NewInt(1337) // Ganache 默认链 ID
	authOptions, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Printf("创建交易选项失败: %v", err)
		return "", errors.New("创建交易选项失败")
	}
	authOptions.Context = context.Background()

	// 获取当前时间戳并将金额转换为 *big.Int
	timestamp := big.NewInt(time.Now().Unix())
	amountBigInt := big.NewInt(amount)

	// 执行智能合约的 RecordTransaction 方法
	tx, err := instance.RecordTransaction(
		authOptions,
		txID,         // 交易 ID
		sender,       // 发送方地址
		receiver,     // 接收方地址
		amountBigInt, // 交易金额
		currencyType, // 货币类型
		timestamp,    // 时间戳
		note,         // 附加备注信息
	)
	if err != nil {
		log.Printf("执行智能合约失败: %v", err)
		return "", errors.New("执行智能合约失败")
	}

	// 打印并返回交易哈希
	txHash := tx.Hash().Hex()
	log.Printf("交易成功提交！交易哈希: %s", txHash)
	return txHash, nil
}
