package utils

import (
	"fmt"
	"regexp"
	"testing"

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
	re := regexp.MustCompile(`\w+\.\w+\.\w+`)
	matchString := re.FindString(jwtToken)
	assert.NotNil(t, matchString)
	assert.Equal(t, jwtToken, matchString)
}
