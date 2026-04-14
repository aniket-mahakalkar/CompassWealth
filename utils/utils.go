package utils

import (
	"encoding/base64"
	"fmt"

	"github.com/charmbracelet/log"
	"golang.org/x/crypto/bcrypt"
)

func ThrowError(err error, msg string) error {
	log.Error(err, msg)
	return fmt.Errorf("%s", msg)
}

func CheckPasswordHash(password, hash string) bool {
	decoded, err := base64.StdEncoding.DecodeString(password)

	if err != nil {
		return false
	}

	err = bcrypt.CompareHashAndPassword([]byte(hash), decoded)
	return err == nil
}

func HashPassword(password string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(password)

	if err != nil {
		return "", err
	}

	bytes, err := bcrypt.GenerateFromPassword(decoded, 14)
	if err != nil {
		log.Warn("error in generating password", err)
	}
	return string(bytes), nil
}
