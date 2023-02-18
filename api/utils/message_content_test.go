package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateMessageContentWithToken(t *testing.T) {
	messageContent := GenerateMessageContentWithToken("token")
	expectMessage := fmt.Sprintf("リクエストありがとうございます。\n\n認証トークンは\n\n%s\nです。", "token")
	assert.Equal(t, expectMessage, messageContent)
}
