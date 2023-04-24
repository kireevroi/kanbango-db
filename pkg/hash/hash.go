package hash

import (
	"fmt"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(pass string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), 14)
	return string(bytes), err
}

func CheckPassword(hash string, pass string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass)) == nil
}

func GenerateUUID() (string, error) {
	x, err := uuid.NewRandom()
	return fmt.Sprint(x), err
}