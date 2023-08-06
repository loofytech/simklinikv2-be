package controllers

import (
	"sim-klinikv2/config"
	"sim-klinikv2/models"
	"sim-klinikv2/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AuthBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func LoginUser(c *fiber.Ctx) error {
	// get body request
	body := new(AuthBody)
	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}
	// query to user table for get one data
	var user models.User
	result := config.DB.First(&user, "email = ?", body.Email)

	// check jika user email tidak ditemukan
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(400).JSON(fiber.Map{"status": "fail", "message": "User tidak ditemukan"})
		}
		return c.Status(500).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	// check user password
	checkPassword := utils.CheckPasswordHash(body.Password, user.Password)

	// check jika user password salah
	if !checkPassword {
		return c.Status(400).JSON(fiber.Map{
			"status":  "failed",
			"message": "Email atau password anda salah",
		})
	}

	// logic jwt

	return c.Status(200).JSON(fiber.Map{
		"status": "success",
		"data":   "login berhasil",
	})
}
