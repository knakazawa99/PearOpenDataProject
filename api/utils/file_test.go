package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateOutPutFileName(t *testing.T) {
	fileName := "test.zip"
	version := "1.0.0"
	expectValue := "test.1.0.0.zip"
	result := GenerateOutPutFileName(fileName, version)
	assert.Equal(t, result, expectValue)
}

func TestGenerateOutPutFileName2(t *testing.T) {
	fileName := "test.zip"
	version := "1.0.10"
	expectValue := "test.1.0.10.zip"
	result := GenerateOutPutFileName(fileName, version)
	assert.Equal(t, result, expectValue)
}
