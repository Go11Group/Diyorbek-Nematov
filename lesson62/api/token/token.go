package token

import (
	"errors"
	"students/models"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	SigningKey = "my_secret_ky"
)

type Claims struct {
	UserId   string
	Username string
	Role     string
	jwt.StandardClaims
}

func GenerateAccessJWT(signUp *models.LoginRequest) (string, error) {

	claims := Claims{
		Username: signUp.Username,
		Role:     signUp.Role,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(30 * time.Minute).Unix(),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return accessToken.SignedString([]byte(SigningKey))
}

func ExtractClaimsAccess(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(SigningKey), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

func ValidateToken(tokenStr string) (bool, error) {
	_, err := ExtractClaimsAccess(tokenStr)
	if err != nil {
		return false, err
	}
	return true, nil
}
