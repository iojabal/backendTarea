package services

import (
	"encoding/base64"

	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(pass string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func EncodeBase64(hash string) string {
	return base64.StdEncoding.EncodeToString([]byte(hash))
}

func DecodeBase64(bs string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(bs)
	if err != nil {
		return "", err
	}
	return string(decoded), nil
}
