package auth

import (
	"errors"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/gin-gonic/gin"
)

var jwtKey = []byte("your_secret_key")

type Claims struct {
	UserID   uint   `json:"user_id"`
	UserName string `json:"username"`
	jwt.StandardClaims
}

// GenerateJWT 生成 JWT 令牌
func GenerateJWT(user *User) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour) // 设置令牌过期时间为 24 小时
	claims := &Claims{
		UserID:   user.UserID,
		UserName: user.UserName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// ValidateJWT 验证 JWT 令牌
func ValidateJWT(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}

// GetUserIDFromToken 从 token 中解析并获取用户 ID
func GetUserIDFromToken(ctx *gin.Context) (uint, error) {
	tokenString := ctx.GetHeader("Authorization")
	claims := &Claims{}

	// 解析 token
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	claims, _ = ValidateJWT(tokenString)

	return claims.UserID, nil
}

// JWTAuthMiddleware JWT 中间件，用于保护需要身份验证的路由
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
			c.Abort()
			return
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		// 验证 JWT
		claims, err := ValidateJWT(tokenString)
		if err != nil {
			log.Println("Token verification failed:", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// 增加日志，检查 user_id 是否正确解析
		log.Println("UserID from JWT:", claims.UserID)
		log.Println("Username from JWT:", claims.UserName)

		// 将 user_id 存储到上下文中
		c.Set("user_id", claims.UserID)
		c.Set("username", claims.UserName)

		c.Next()
	}
}
