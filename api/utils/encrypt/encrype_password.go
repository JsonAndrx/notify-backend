package encrypt

import (
	"golang.org/x/crypto/bcrypt"
	"notify-backend/api/utils/debug"

)

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		debug.LogError(err)
		return "", err
	}
	return string(hash), nil
}