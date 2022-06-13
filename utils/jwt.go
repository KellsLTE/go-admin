package utils

import (
	"time"

	"github.com/KellsLTE/go-admin/config"
	"github.com/dgrijalva/jwt-go"
)

func GenerateJwt(issuer string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer: issuer,
		ExpiresAt: time.Now().Add(time.Hour * 12).Unix(), // 1/2 day
	})

	SecretKey := config.Env("JWT_SECRET_KEY")

	return claims.SignedString([]byte(SecretKey))
}

func ParseJwt(cookie string) (string, error) {
	SecretKey := config.Env("JWT_SECRET_KEY")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error){
		return []byte(SecretKey), nil
	}) 

	if err != nil || !token.Valid {
		return "", err
	}

	claims := token.Claims.(*jwt.StandardClaims)

	return claims.Issuer, nil
}