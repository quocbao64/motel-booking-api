package pkg

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	passwordHashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(passwordHashed), err
}

func ComparePassword(password string, passwordHashed string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(passwordHashed), []byte(password))
	return err == nil
}
