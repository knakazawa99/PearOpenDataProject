package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(targetString string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"targetString": targetString,
		"nbf":          time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})
	t, _ := token.SignedString([]byte("SECRET_KEY"))
	return t
}
