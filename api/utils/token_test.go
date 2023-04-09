package utils

import (
	"fmt"
	"testing"

	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
)

func TestGenerateToken(t *testing.T) {
	token := GenerateToken("test")
	expectType := "test"
	fmt.Println(token)
	assert.IsType(t, expectType, token)
}

func TestGenerateJWT(t *testing.T) {
	id := "122"
	jwtToken := GenerateJWT(id)
	parsedToken, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		return []byte("SECRET_KEY"), nil
	})
	assert.Nil(t, err)
	assert.True(t, parsedToken.Valid)

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	assert.True(t, ok)
	assert.True(t, parsedToken.Valid)

	key, ok := claims["id"].(string)
	assert.True(t, ok)
	assert.Equal(t, "122", key)

}
