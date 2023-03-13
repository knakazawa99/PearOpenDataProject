package utils

import (
	"fmt"
	"time"
)

func GenerateMessageContentWithToken(token string) string {
	now := time.Now()
	return fmt.Sprintf("リクエストありがとうございます。\n\n"+
		"認証トークンは\n\n%s\n\nです。\n\n\n"+
		"Copyright Yamazaki Lab. Niigata Univ, All rights reserved.%v.", token, now.Year())
}
