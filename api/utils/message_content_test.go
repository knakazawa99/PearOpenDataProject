package utils

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGenerateMessageContentWithToken(t *testing.T) {
	now := time.Now()
	messageContent := GenerateMessageContentWithToken("token")
	expectMessage := fmt.Sprintf("リクエストありがとうございます。\n\n認証トークンは\n\n%s\n\nです。\n\n\nCopyright Yamazaki Lab. Niigata Univ, All rights reserved.%v.", "token", now.Year())
	assert.Equal(t, expectMessage, messageContent)
}
