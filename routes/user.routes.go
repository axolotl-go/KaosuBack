package routes

import (
	"github.com/axolotl-go/KaosuBack/db"
	"github.com/axolotl-go/KaosuBack/models"
	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {
	var user models.Users

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	if err := db.DB.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create user",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(user)

}

func FindUsers(c *fiber.Ctx) error {
	var users []models.Users

	if err := db.DB.Find(&users).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(users)

}

func FindUser(c *fiber.Ctx) error {
	var user models.Users
	id := c.Params("id")

	if err := db.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(user)
}
