package utils

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(targetString string) string {
	rand.Seed(time.Now().UnixNano())
	// 10桁
	randomNum := rand.Intn(9000000000) + 1000000000
	return fmt.Sprintf("%d", randomNum)
}
func GenerateJWT(id string) string {
	claims := jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	// ヘッダーとペイロードの生成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte("SECRET_KEY"))
	return tokenString
}
