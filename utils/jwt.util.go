package utils

import "github.com/golang-jwt/jwt/v5"


var scret_key = "SCRERT_KEY"
func GenerateToken(claims *jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	webtoken, error := token.SignedString([]byte(scret_key))

	if error != nil {
		return "", error
	}
	return webtoken, nil
}