package util

import (
	"log"

	"github.com/duyk16/secure-rest-api/config"

	jwt "github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type Token struct {
	ID    primitive.ObjectID `json:"id"`
	Email string             `json:"email"`
	jwt.StandardClaims
}

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

func GenerateToken(userId primitive.ObjectID, email string) string {
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), Token{
		ID:    userId,
		Email: email,
	})
	tokenString, _ := token.SignedString([]byte(config.ServerConfig.JWTKey))
	return tokenString
}
