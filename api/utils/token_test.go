package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateToken(t *testing.T) {
	token := GenerateToken("test")
	expectType := "test"
	fmt.Println(token)
	assert.IsType(t, expectType, token)
}
