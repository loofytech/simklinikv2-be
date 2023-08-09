package middleware

import (
	"sim-klinikv2/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// type AuthMidlleware() fiber.Handler {
// 	return func(c fiber.Ctx) error {
// 		tokenString := c.GetRespHeader("Authorization")
// 		if tokenString == "" {
// 			c.Status(401).JSON(fiber.Map{"error": "request tidak berisi akses token"})
// 		}
// 	}
// }
// func NewAuthMiddleware(secret string) fiber.Handler {
// 	return jwtware.New(jwtware.Config{
// 		SigningKey: []byte(secret),
// 	})
// }

func Auth(c *fiber.Ctx) error {
	tokenString := c.Get("Authorization")
	if tokenString == "" {
		return c.Status(401).JSON(fiber.Map{
			"message": "Unauthenticated 1",
		})
	}

	token := strings.Split(tokenString, " ")[1]

	_, err := utils.ValidateToken(token)
	// claims, err :=utils.DecodeToken()
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"message": "Unauthenticated 2",
		})
	}

	// decodeToken, err := utils.DecodeToken(token)

	// c.Locals("userInfo", claims)
	// c.Locals("role", claims)

	return c.Next()
}
