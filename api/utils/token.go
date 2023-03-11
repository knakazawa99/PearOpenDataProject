package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateToken(targetString string) string {
	rand.Seed(time.Now().UnixNano())
	// 10桁
	randomNum := rand.Intn(9000000000) + 1000000000
	return fmt.Sprintf("%d", randomNum)
}
