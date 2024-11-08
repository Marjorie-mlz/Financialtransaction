package auth

import (
	"time"

	"gorm.io/gorm"
)

// 确保导入包含 Account 模型的包
// 如果 Account 定义在同一个包中（如 auth 包），则不需要额外导入

// UserRepository 定义用户和交易相关的操作接口
type UserRepository interface {
	Save(user *User) error
	FindByUsername(username string) (*User, error)
	CreateAccount(userID uint, ethereumAddress string) error
	CreateTransaction(tx *TransactionRecord) error
	GetTransactionByID(txID string) (*TransactionRecord, error)
	SaveTransactionStatus(txID, status string) error
	GetTransactionStatus(txID string) (string, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

// 保存用户到数据库
func (r *userRepository) Save(user *User) error {
	return r.db.Create(user).Error
}

// 根据用户名查找用户
func (r *userRepository) FindByUsername(username string) (*User, error) {
	var user User
	result := r.db.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// 创建用户账户
func (r *userRepository) CreateAccount(userID uint, ethereumAddress string) error {
	// 创建 Account 实例，并设置账户相关信息
	account := Account{
		UserID:          userID,
		EthereumAddress: ethereumAddress,
		AccountType:     "personal", // 假设所有账户类型为 "personal"
		AccountStatus:   "active",   // 假设新建账户状态为 "active"
	}
	return r.db.Create(&account).Error
}

// GetUserByID 根据用户 ID 获取用户信息
func (s *AuthService) GetUserByID(userID uint) (*User, error) {
	var user User
	if err := s.db.First(&user, userID).Error; err != nil { // 使用 s.db 而不是 db
		return nil, err
	}
	return &user, nil
}

// 添加账户余额更新
func (r *userRepository) UpdateAccountBalance(userID uint, newBalance float64) error {
	return r.db.Model(&Account{}).Where("user_id = ?", userID).Update("balance", newBalance).Error
}

// 创建交易记录并保存到 transactionrecord 表
func (r *userRepository) CreateTransaction(tx *TransactionRecord) error {
	return r.db.Create(tx).Error
}

// 根据交易 ID 查询交易记录
func (r *userRepository) GetTransactionByID(txID string) (*TransactionRecord, error) {
	var tx TransactionRecord
	if err := r.db.Where("tx_id = ?", txID).First(&tx).Error; err != nil {
		return nil, err
	}
	return &tx, nil
}

// 保存交易状态到 transactionstatus 表
func (r *userRepository) SaveTransactionStatus(txID, status string) error {
	return r.db.Create(&TransactionStatus{
		TxID:      txID,
		Status:    status,
		Timestamp: time.Now(),
	}).Error
}

// 根据交易 ID 获取交易状态
func (r *userRepository) GetTransactionStatus(txID string) (string, error) {
	var status TransactionStatus
	if err := r.db.Where("tx_id = ?", txID).First(&status).Error; err != nil {
		return "", err
	}
	return status.Status, nil
}
