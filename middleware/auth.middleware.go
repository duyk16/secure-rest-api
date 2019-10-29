package middleware

import (
	"net/http"
	"strings"

	"github.com/duyk16/secure-rest-api/config"
	u "github.com/duyk16/secure-rest-api/util"

	"github.com/dgrijalva/jwt-go"
)

func JwtAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		noAuthRoutes := []string{
			"/api/auth/login",
			"/api/auth/register",
		}

		//check if request does not need authentication, serve the request if it doesn't need it
		for _, value := range noAuthRoutes {
			if value == r.URL.Path {
				next.ServeHTTP(w, r)
				return
			}
		}

		tokenHeader := r.Header.Get("Authorization")
		if tokenHeader == "" {
			u.JSON(w, 400, u.T{
				"status":  "error",
				"message": "Missing token",
			})
			return
		}

		// `Bearer {token-body}`
		splitted := strings.Split(tokenHeader, " ")
		if len(splitted) != 2 {
			u.JSON(w, 400, u.T{
				"status":  "error",
				"message": "Invalid/Malformed auth token",
			})
			return
		}

		tokenString := splitted[1]
		token := u.Token{}

		tokenData, err := jwt.ParseWithClaims(tokenString, &token, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.ServerConfig.JWTKey), nil
		})

		if err != nil {
			u.JSON(w, http.StatusForbidden, u.T{
				"status":  "error",
				"message": "Malformed authentication token",
			})
			return
		}

		if !tokenData.Valid {
			u.JSON(w, http.StatusForbidden, u.T{
				"status":  "error",
				"message": "Token is not valid.",
			})
			return
		}

		next.ServeHTTP(w, r)
	})
}
