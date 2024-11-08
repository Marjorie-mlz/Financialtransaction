package auth

import (
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User 模型，映射到数据库中的 users 表
type User struct {
	UserID       uint      `json:"user_id" gorm:"column:UserID;primaryKey;autoIncrement"`
	UserName     string    `json:"username" gorm:"column:UserName;unique;not null"`    // 用户名，必须唯一
	PasswordHash string    `json:"password_hash" gorm:"column:password_hash;not null"` // 存储哈希后的密码
	Password     string    `json:"password" gorm:"-"`                                  // 临时存储明文密码，不存储在数据库中
	Email        string    `json:"email" gorm:"column:Email;unique;not null"`          // 邮箱，必须唯一
	CreatedAt    time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"column:updated_at"`
	Accounts     []Account `gorm:"foreignKey:UserID"` // 与 Account 模型的关联，外键为 UserID
}

// Account 模型，映射到数据库中的 account 表
type Account struct {
	AccountID       uint      `json:"account_id" gorm:"column:AccountID;primaryKey;autoIncrement"`
	UserID          uint      `json:"user_id" gorm:"column:UserID;not null"`                            // 外键，关联到 User
	EthereumAddress string    `json:"ethereum_address" gorm:"column:EthereumAddress;size:255;not null"` // 以太坊地址
	Balance         float64   `gorm:"default:0"`
	AccountType     string    `json:"account_type" gorm:"column:AccountType;size:50;not null"`     // 账户类型
	AccountStatus   string    `json:"account_status" gorm:"column:AccountStatus;size:50;not null"` // 账户状态
	CreatedAt       time.Time `json:"created_at" gorm:"column:CreatedAt"`                          // 账户创建时间
	UpdatedAt       time.Time `json:"updated_at" gorm:"column:UpdatedAt"`                          // 账户更新时间
}

func (a *Account) TableName() string {
	return "account"
}

// TransactionRecord 模型，映射到 transactionrecord 表
type TransactionRecord struct {
	TxID            string    `gorm:"primaryKey;type:varchar(255)" json:"tx_id"`
	SenderAccount   string    `gorm:"column:SenderAccount" json:"sender_account"`
	ReceiverAccount string    `gorm:"column:ReceiverAccount" json:"receiver_account"`
	Amount          float64   `gorm:"type:decimal(18,8)" json:"amount"`
	CurrencyType    string    `gorm:"type:varchar(10)" json:"currency_type"`
	CurrencyUnit    string    `gorm:"type:varchar(10)" json:"currency_unit"`
	TxHash          string    `gorm:"type:varchar(255)" json:"tx_hash"`
	BlockHeight     int       `gorm:"column:BlockHeight" json:"block_height"`
	Note            string    `gorm:"type:text" json:"note,omitempty"`
	Status          string    `gorm:"type:varchar(50)" json:"status"`
	Timestamp       time.Time `gorm:"column:Timestamp" json:"timestamp"`
	CreatedAt       time.Time `gorm:"column:CreatedAt" json:"created_at"`
	UserID          uint      `json:"user_id" gorm:"column:UserID;"`
}

// TransactionStatus 模型，映射到 transactionstatus 表
type TransactionStatus struct {
	TxID      string    `gorm:"primaryKey;type:varchar(255)" json:"tx_id"`
	Status    string    `gorm:"type:varchar(50)" json:"status"`
	Timestamp time.Time `gorm:"column:Timestamp" json:"timestamp"`
}

// TransactionRequest 模型，用于 API 请求
type TransactionRequest struct {
	SenderAccount   string  `json:"sender" from:"sender"` // 发送方账户
	ReceiverAccount string  `json:"receiver" from:"receiver"`
	Amount          float64 `json:"amount" from:"amount"`
	CurrencyType    string  `json:"currency" from:"currency"`
	CurrencyUnit    string  `json:"currency_unit" from:"currency_unit"` // 添加此字段
	Remarks         string  `json:"remarks" form:"remarks"`
}

// BeforeCreate GORM 的模型钩子，创建用户模型前调用
func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	// 检查密码是否已经被哈希，避免重复哈希
	if !IsHashed(user.Password) {
		// 如果密码没有哈希，则对密码进行哈希处理
		hashedPassword, err := HashPassword(user.Password)
		if err != nil {
			log.Println("Error hashing password:", err)
			return err
		}
		user.PasswordHash = hashedPassword
		// 打印生成的哈希密码，用于调试
		log.Println("Generated Hash Password (Register):", user.PasswordHash)
	}
	return nil
}

// IsHashed 判断密码是否已经被哈希
func IsHashed(password string) bool {
	// 检查哈希格式 (bcrypt 哈希格式以 $2a$ 开头)
	return len(password) > 0 && password[:4] == "$2a$"
}

// HashPassword 使用 bcrypt 对密码进行哈希
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// BeforeCreate 钩子，在创建交易记录前自动设置时间戳
func (tx *TransactionRecord) BeforeCreate(_ *gorm.DB) (err error) {
	if tx.Timestamp.IsZero() { // 如果 Timestamp 未设置，自动设置为当前时间
		tx.Timestamp = time.Now()
	}
	return nil
}
