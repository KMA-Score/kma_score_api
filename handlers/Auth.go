package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"log"
)

type AuthTokenType struct {
	Token string `json:"token"`
}

func AuthToken(c *fiber.Ctx) error {
	p := new(AuthTokenType)

	if err := c.BodyParser(p); err != nil {
		log.Println(err)
	}

	token, err := jwt.Parse(p.Token, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return hmacSampleSecret, nil
	})
	if err != nil {
		panic(err)
	}
	claims := token.Claims.(jwt.MapClaims)

}
