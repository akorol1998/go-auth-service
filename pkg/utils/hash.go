package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(pass string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(pass), 5)
	return string(bytes)
}

func CheckPasswordHash(hashPass string, rawPass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPass), []byte(rawPass))

	return err == nil
}
