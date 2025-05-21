package utils

import (
	"api-blog-go/internal/config"
	"api-blog-go/internal/models"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey []byte // Nên đặt trong biến môi trường
func init() {
	config.LoadEnv()
	secretKey = []byte(config.GetEnv("JWT_KEY", ""))
}
func GenerateToken(user *models.User) (string, error) {
	// Đảm bảo user.ID là uint
	claims := jwt.MapClaims{
		"user_id":  float64(user.ID), // Chuyển đổi uint sang float64
		"email":    user.Email,
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

func ValidateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Kiểm tra phương thức ký
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	// Kiểm tra token có hợp lệ không
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return token, nil
}
