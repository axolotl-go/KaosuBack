package routes

import (
	"github.com/axolotl-go/KaosuBack/db"
	"github.com/axolotl-go/KaosuBack/models"
	"github.com/gofiber/fiber/v2"
)

func CreateProduct(c *fiber.Ctx) error {
	var products models.Products

	if err := c.BodyParser(&products); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	if err := db.DB.Create(&products).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create user",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(products)
}

func GetProducts(c *fiber.Ctx) error {
	var products []models.Products

	if err := db.DB.Find(&products).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(products)
}

func GetProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	var product models.Products

	if err := db.DB.Where("id = ?", id).First(&product).Error; err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Product not found")
	}

	return c.JSON(product)
}
