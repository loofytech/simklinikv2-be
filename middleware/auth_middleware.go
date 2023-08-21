package middleware

import (
	"fmt"
	"sim-klinikv2/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// type user interface{}

func Auth(c *fiber.Ctx) error {
	tokenString := c.Get("Authorization")
	if tokenString == "" {
		return c.Status(401).JSON(fiber.Map{
			"message": "Unauthenticated 1",
		})
	}

	token := strings.Split(tokenString, " ")[1]

	_, err := utils.ValidateToken(token)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"message": "Unauthenticated 2",
		})
	}

	decodeToken, err := utils.DecodeToken(token)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"message": "Unauthenticated 2",
		})
	}
	c.Locals("userInfo", decodeToken)

	return c.Next()
}

func CheckUserRoleAdmin(c *fiber.Ctx) error {
	info := c.Locals("userInfo")
	fmt.Print(info)

	return c.Next()
}
