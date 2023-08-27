package controllers

import (
	"sim-klinikv2/config"
	"sim-klinikv2/models"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateRecipe(c *fiber.Ctx) error {
	var payload *models.CreateRecipeSchema

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	errors := models.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)

	}

	now := time.Now()
	newRecipe := models.Recipe{
		RegistrationId: payload.RegistrationId,
		MedicineId:     payload.MedicineId,
		Quantity:       payload.Quantity,
		SubTotal:       payload.SubTotal,
		CreatedAt:      now,
		UpdatedAt:      now,
	}

	result := config.DB.Create(&newRecipe)

	if result.Error != nil && strings.Contains(result.Error.Error(), "duplicate key value violates unique") {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"status": "fail", "message": "Recipe already exist"})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"data":   newRecipe,
	})
}

func FindRecipe(c *fiber.Ctx) error {
	var page = c.Query("page", "1")
	var limit = c.Query("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var recipe []models.Recipe
	results := config.DB.Limit(intLimit).Offset(offset).Find(&recipe)
	if results.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": results.Error})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "results": len(recipe), "data": recipe})
}

func UpdateRecipe(c *fiber.Ctx) error {
	recipeId := c.Params("recipeId")

	var payload *models.UpdateRecipeSchema

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	var recipe models.Recipe
	result := config.DB.First(&recipe, "id = ?", recipeId)
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No Recipe with that Id exists"})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	updates := make(map[string]interface{})
	if payload.Quantity != 0 {
		updates["quantity"] = payload.Quantity
	}

	if payload.SubTotal != 0 {
		updates["sub_total"] = payload.SubTotal
	}

	if payload.RegistrationId != 0 {
		updates["registration_id"] = payload.RegistrationId
	}

	if payload.MedicineId != 0 {
		updates["medicine_id"] = payload.MedicineId
	}

	updates["updated_at"] = time.Now()

	config.DB.Model(&recipe).Updates(updates)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": recipe})
}

func FindRecipeById(c *fiber.Ctx) error {
	recipeId := c.Params("recipeId")

	var recipe models.Recipe
	result := config.DB.First(&recipe, "id = ?", recipeId)
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No Recipe with that Id exists"})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"Recipe": recipe}})
}

func RecipeDelete(c *fiber.Ctx) error {
	recipeId := c.Params("recipeId")

	result := config.DB.Delete(&models.Recipe{}, "id = ?", recipeId)

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No Recipe with that Id exists"})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error})
	}

	return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error})
}
