package middleware

import (
	"fmt"
	"melody_cure/model"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type JwtClient struct {
	SecretKey string
}

// 生成token
func (jc *JwtClient) GenerateToken(id string) (string, error) {
	claims := model.Claims{
		UserId: id, // 使用传入的用户 ID
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(0, 1, 0).Unix(), // 过期时间设置为 1 个月后
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jc.SecretKey)) // 替换为你的密钥
}

// 创建一个中间件来验证 JWT 并检查用户角色：
func (jc *JwtClient) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(401, gin.H{"message": "未认证"})
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"]) 
			}
			return []byte(jc.SecretKey), nil
		})

		if err != nil {
			c.JSON(401, gin.H{"message": "Invalid token", "error": err.Error()})
			c.Abort()
			return
		}

		if !token.Valid {
			c.JSON(401, gin.H{"message": "Invalid or expired token"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			c.JSON(401, gin.H{"message": "未认证"})
			c.Abort()
			return
		}

		// 获取 user_id 为字符串类型
		userId, ok := claims["user_id"].(string)
		if !ok {
			c.JSON(401, gin.H{"message": "无法获取id"})
			c.Abort()
			return
		}

		// 存储在上下文中
		c.Set("user_id", userId)

		c.Next()
	}
}

func (jc *JwtClient) ParseToken(tokenString string) (string, error) {
	if tokenString == "" {
		return "", fmt.Errorf("token is empty")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"]) 
		}
		return []byte(jc.SecretKey), nil // 使用 SecretKey
	})

	if err != nil {
		return "", fmt.Errorf("token is invalid")
	}

	if !token.Valid {
		return "", fmt.Errorf("token is invalid")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", fmt.Errorf("未认证")
	}

	// 获取 user_id 为字符串类型
	userId, ok := claims["user_id"].(string)
	if !ok {
		return "", fmt.Errorf("无法获取user_id")
	}

	return userId, nil
}
