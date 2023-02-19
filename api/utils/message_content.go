package utils

import (
	"fmt"
)

func GenerateMessageContentWithToken(token string) string {
	return fmt.Sprintf("リクエストありがとうございます。\n\n認証トークンは\n\n%s\nです。", token)
}
