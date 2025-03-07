package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Email string
	Role  string
	jwt.RegisteredClaims
}

func (a *Authentication) GenerateJWT(email, role string) (string, error) {
	// Set the claims
	claims := Claims{
		Role:  role,
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 1)), // Token expires in 1 hour
			Issuer:    "MyProject",
		},
	}

	// Create the token using the claims and the signing key
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate and return the signed token
	tokenString, err := token.SignedString([]byte(a.SignInKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (a *Authentication) ValidateToken(token string) (Claims, error) {
	// Parse the token with the registered claims.
	var claims Claims
	tkn, err := jwt.ParseWithClaims(token, &claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(a.SignInKey), nil
	})
	if err != nil {
		return Claims{}, fmt.Errorf("error in parsing the token : %w", err)
	}

	// checking if the token is valid or not
	if !tkn.Valid {
		return Claims{}, errors.New("token in not valid")
	}

	return claims, nil

}

