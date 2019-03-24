package router

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

// Result of token verification
type Result struct {
	Username string
	OK       bool
}

// Claims - jwt claims
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// Ideally saved in secure store (s3) and injected into environment variable of servers
var jwtKey = []byte("a62f2225bf70bfaccbc7f1ef2a397836717377de")

// VerifyToken - verifies the jwt token
func VerifyToken(tokenString string) (Result, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if !token.Valid {
		return Result{
			OK: false,
		}, nil
	}
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return Result{
				OK: false,
			}, nil
		}

		return Result{OK: false}, errors.Wrap(err, "Unable to verify JWT")
	}

	return Result{
		OK:       true,
		Username: claims.Username,
	}, nil
}
