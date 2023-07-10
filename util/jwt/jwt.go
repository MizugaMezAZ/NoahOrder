package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	UserID int64 `json:"user_id"`
	jwt.StandardClaims
}

// ExpirationTime ...
var ExpirationTime = 24 * time.Hour

var JwtSecret = []byte("bibibublbi9121")

func Gentoken(uid int64) (string, error) {
	expirationTime := time.Now().Add(ExpirationTime)

	claims := &Claims{
		UserID: uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "bibibububi",
			Subject:   "user token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
