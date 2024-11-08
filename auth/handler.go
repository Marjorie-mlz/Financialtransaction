package auth

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/seencxy/lsr/common"

	"github.com/gin-gonic/gin"
)

// CORS 中间件，设置允许的跨域请求头
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		// 对于OPTIONS请求，直接返回200
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		c.Next()
	}
}

type AuthHandler struct {
	service *AuthService
}

func NewAuthHandler(service *AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

// Register 方法，用户注册
func (h *AuthHandler) Register(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// 传递指针类型给 Register 方法
	if err := h.service.Register(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

// Login 方法，用户登录
func (h *AuthHandler) Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	user, err := h.service.Authenticate(input.Email, input.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// 生成 JWT token
	token, err := h.service.GenerateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// 返回 token 给前端
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
	})
}

// GetProfile 方法，获取用户信息, 使用 nonce
func (h *AuthHandler) GetProfile(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found"})
		return
	}

	user, err := h.service.GetUserByID(userID.(uint))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	// 从上下文中获取 nonce
	nonce, _ := c.Get("nonce")

	c.JSON(http.StatusOK, gin.H{
		"username": user.UserName,
		"email":    user.Email,
		"nonce":    nonce,
	})
}

// ChangePassword 方法，修改用户密码
func (h *AuthHandler) ChangePassword(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		log.Println("User ID not found in context")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found"})
		return
	}
	log.Println("User ID from context:", userID)

	var input struct {
		OldPassword string `json:"old_password"`
		NewPassword string `json:"new_password"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	user, err := h.service.GetUserByID(userID.(uint))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// 验证旧密码是否正确
	if err := h.service.VerifyPassword(user.PasswordHash, input.OldPassword); err != nil {
		log.Println("Old password verification failed for user:", userID)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect old password"})
		return
	}

	// 更新密码
	if err := h.service.UpdatePassword(userID.(uint), input.NewPassword); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password changed successfully. Please log in again"})
}

// JWTMiddleware 验证JWT的中间件
func (h *AuthHandler) JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := h.service.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// 将用户信息存入上下文，供后续处理使用
		c.Set("user", claims)
		c.Next()
	}
}

// GetBalance 方法：获取指定以太坊地址的余额
func (h *AuthHandler) GetBalance(c *gin.Context) {
	userId, err := GetUserIDFromToken(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ethereum address is required"})
		return
	}

	ethereumAddress, err := SelectUserById(int(userId), h.service.db)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ethereum address is required"})
		return
	}
	if ethereumAddress == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ethereum address is required"})
		return
	}

	// 调用服务层的方法获取余额
	balance, err := h.service.GetBalance(ethereumAddress)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get balance"})
		return
	}

	// 返回账户余额
	c.JSON(http.StatusOK, gin.H{"balance": balance})
}

// SubmitTransaction 方法：将交易记录上链并存储数据库
func (h *AuthHandler) SubmitTransaction(c *gin.Context) {
	var txRequest TransactionRequest

	// 解析 JSON 请求体
	if err := c.ShouldBindJSON(&txRequest); err != nil {
		log.Printf("Error binding JSON: %v", err) // 打印具体错误
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	mal, _ := json.Marshal(&txRequest)
	fmt.Println("mal:", string(mal))
	fmt.Println("11111:")
	// 从上下文中获取用户 ID
	userId, _ := GetUserIDFromToken(c)
	// 确保 userID 与交易发送者账户一致（假设账户ID为字符串类型地址）
	// txRequest.SenderAccount = string(userID.(uint))
	// sa,_:=strconv.Atoi(req.SenderAccount)
	log.Printf("Sender Address: %s", txRequest.SenderAccount) // 打印发送者地址
	fmt.Println("333333:")
	// 打印接收到的字段，确认传输无误
	log.Printf("Received TransactionRequest: %+v", txRequest)

	// 调用服务层方法进行上链并存储数据库
	txHash, err := h.service.CallSmartContractTransfer(&txRequest)
	if err != nil {
		log.Printf("Error recording transaction on-chain: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to record transaction on blockchain"})
		return
	}
	// 打印成功信息
	log.Printf("Transaction recorded successfully with hash: %s", txHash)
	fmt.Println("44444444:")
	// 创建交易记录，并将其存入数据库
	transaction := TransactionRecord{
		TxID:            common.GenerateRandomString(10),
		SenderAccount:   txRequest.SenderAccount,
		ReceiverAccount: txRequest.ReceiverAccount,
		Amount:          txRequest.Amount,
		CurrencyType:    txRequest.CurrencyType,
		Note:            txRequest.Remarks,
		Status:          "Completed",
		TxHash:          txHash,
		Timestamp:       time.Now(),
		UserID:          userId,
	}
	fmt.Println("Transaction:", transaction)
	// 将交易记录存入数据库
	if err := h.service.db.Debug().Create(&transaction).Error; err != nil {
		log.Printf("Error saving transaction to database: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save transaction to database"})
		return
	}

	// 返回成功信息，包括交易哈希
	c.JSON(http.StatusOK, gin.H{
		"message":         "Transaction submitted successfully",
		"transactionHash": txHash,
	})
}

// 查询交易记录API
func (h *AuthHandler) GetTransactionByTxID(c *gin.Context) {
	txID := c.Param("txID")

	var transaction TransactionRecord
	if err := h.service.db.Where("tx_id = ?", txID).First(&transaction).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
		return
	}

	c.JSON(http.StatusOK, transaction)
}

// GetTransactions 方法：获取用户最近的交易记录
func (h *AuthHandler) GetTransactions(c *gin.Context) {
	userId, err := GetUserIDFromToken(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ethereum address is required"})
		return
	}
	// 查询用户最近的交易记录
	transactions, err := h.service.GetRecentTransactionsByAddress(int(userId))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch transactions"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"transactions": transactions})
}
