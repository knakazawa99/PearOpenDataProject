package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestPasswordEncrypt(t *testing.T) {
	password := "password"
	result, err := PasswordEncrypt(password)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.NotEqual(t, result, password)
}

func TestCheckHashPassword(t *testing.T) {
	password := "password"
	encryptedPassword, _ := PasswordEncrypt(password)
	err := CheckHashPassword(encryptedPassword, password)
	assert.Nil(t, err)
}

func TestCheckHashPasswordFail(t *testing.T) {
	password := "password"
	encryptedPassword, _ := PasswordEncrypt(password)
	err := CheckHashPassword(encryptedPassword, "password1")
	assert.Equal(t, err.Error(), bcrypt.ErrMismatchedHashAndPassword.Error())
	assert.NotNil(t, err)
}
