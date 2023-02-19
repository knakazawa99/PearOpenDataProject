package utils

import (
	"fmt"
	"strings"
)

func GenerateOutPutFileName(fileName string, version string) string {
	parts := strings.Split(fileName, ".")
	ext := parts[len(parts)-1]
	name := strings.Join(parts[:len(parts)-1], ".")
	return fmt.Sprintf("%s.%s.%s", name, version, ext)
}
