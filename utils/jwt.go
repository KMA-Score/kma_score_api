package utils

import (
	"github.com/golang-jwt/jwt"
	"strings"
)

func JwtVerifyToken(token string, pemKey string) (bool, error) {
	key, err := jwt.ParseRSAPublicKeyFromPEM([]byte(pemKey))

	if err != nil {
		return false, err
	}

	parts := strings.Split(token, ".")

	err = jwt.SigningMethodRS256.Verify(strings.Join(parts[0:2], "."), parts[2], key)

	if err != nil {
		return false, nil
	}
	return true, nil
}

func JWTConstructToken(claimsToCreate jwt.Claims, pemKey string) (string, error) {
	key, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(pemKey))

	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claimsToCreate)

	tokenString, err := token.SignedString(key)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
