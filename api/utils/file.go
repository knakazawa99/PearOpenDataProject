package utils

import (
	"fmt"
	"strings"
)

func GenerateOutPutFileName(fileName string, version string) string {
	parts := strings.Split(fileName, ".")
	ext := parts[len(parts)-1]
	fileNameParts := strings.Split(parts[0], "/")
	name := fileNameParts[len(fileNameParts)-1]
	return fmt.Sprintf("%s.%s.%s", name, version, ext)
}
