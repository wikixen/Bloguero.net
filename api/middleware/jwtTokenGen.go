package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/wikixen/blogapp/config"
)


var env = config.GetConfig()

func GenToken(username string) (string, error){
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"username": username,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenStr, err := token.SignedString(env.Secret)
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func verifyToken(tokenStr string) error {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return env.Secret, nil
	})
	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("Invalid token")
	}

	return nil
}

func AccessHandler(next http.Handler)  http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type","application/json")
		tokenStr := r.Header.Get("Authorization")
		if tokenStr == "" {
			http.Error(w,"Token not authorized",http.StatusUnauthorized)
			return
		}
		tokenStr = tokenStr[len("Bearer "):]

		err := verifyToken(tokenStr)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w,r)	
	})
}