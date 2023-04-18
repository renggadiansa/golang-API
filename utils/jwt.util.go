package utils

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

var scret_key = "SCRERT_KEY"

func GenerateToken(claims *jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	webtoken, error := token.SignedString([]byte(scret_key))

	if error != nil {
		return "", error
	}
	return webtoken, nil
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	tokenJwt, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		_, isValid := t.Method.(*jwt.SigningMethodHMAC)
		if !isValid {
			return nil, fmt.Errorf("invalid token: %v", t.Header["alg"])
		}

		return []byte(scret_key), nil
	})

	if err != nil {
		return nil, err
	}

	return tokenJwt, nil
}

func DecodeToken(tokenString string) (jwt.MapClaims, error) {
	token, err := VerifyToken(tokenString)

	if err != nil {
		return nil, err
	}

	claims, isOK := token.Claims.(jwt.MapClaims)
	// fmt.Println("chegoo agr", token.Valid)
	if isOK && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
