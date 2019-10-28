package util

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashAndSaltPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		log.Println("Hash password fail", err)
	}
	return string(hash)
}

func ComparePasswords(hashedPwd string, pwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(pwd))
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
