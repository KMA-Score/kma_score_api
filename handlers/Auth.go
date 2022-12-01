package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/mitchellh/mapstructure"
	"kma_score_api/utils"
	"log"
	"os"
	"time"
)

type AuthTokenType struct {
	Token string `json:"token"`
}

type TokenBody struct {
	Email string
	Exp   int64
	Iat   int64
	Id    string
	Name  string
	MsExp int64
}

func AuthToken(c *fiber.Ctx) error {
	p := new(AuthTokenType)

	if err := c.BodyParser(p); err != nil {
		log.Println(err)
		return c.Status(400).JSON(utils.ApiResponse(400, "Body parse failed", ""))
	}

	if len(p.Token) == 0 {
		return c.Status(400).JSON(utils.ApiResponse(400, "Missing Token param", nil))
	}

	isValid, err := utils.JwtVerifyToken(p.Token, os.Getenv("JWT_CLIENT_PUBLIC_KEY"))

	if err != nil {
		log.Println(err)
		return c.Status(500).JSON(utils.ApiResponse(500, "Error in validate token", nil))
	}

	if isValid == false {
		return c.Status(400).JSON(utils.ApiResponse(400, "Token is invalid", nil))
	}

	// TODO: Figure out how to Parse and Validate in one step. Make it faster
	//rsaPrivate, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(os.Getenv("JWT_CLIENT_PUBLIC_KEY")))

	claims := jwt.MapClaims{}

	_, _, err = new(jwt.Parser).ParseUnverified(p.Token, claims)

	if err != nil {
		fmt.Println(err)
		return c.Status(500).JSON(utils.ApiResponse(500, "Parse token failed", nil))
	}

	// Dirty hack convert map to struct
	var tokenStruct TokenBody
	err = mapstructure.Decode(claims, &tokenStruct)

	if err != nil {
		fmt.Println(err)
		return c.Status(500).JSON(utils.ApiResponse(500, "Parse token struct failed", nil))
	}

	// Check token valid exp
	if tokenStruct.MsExp < time.Now().UTC().Unix() {
		return c.Status(403).JSON(utils.ApiResponse(403, "Token expired", nil))
	}

	newToken, err := utils.JWTConstructToken(claims, os.Getenv("JWT_SERVER_PRIVATE_KEY"))

	if err != nil {
		fmt.Println(err)
		return c.Status(500).JSON(utils.ApiResponse(500, "Generate token failed", nil))
	}

	rsp := fiber.Map{
		"token": newToken,
	}

	return c.Status(200).JSON(utils.ApiResponse(200, "Success", rsp))
}
