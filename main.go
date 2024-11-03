package main

import (
	"github.com/axolotl-go/KaosuBack/db"
	"github.com/axolotl-go/KaosuBack/models"
	"github.com/axolotl-go/KaosuBack/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	db.Dbconnection()
	db.DB.AutoMigrate(models.Products{})
	db.DB.AutoMigrate(models.Users{})

	app := fiber.New()
	app.Use(cors.New(cors.ConfigDefault))

	// User
	app.Post("/user", routes.CreateUser)
	app.Get("/user/:id", routes.FindUser)
	app.Get("/users", routes.FindUsers)

	// Products
	app.Get("/products", routes.GetProducts)
	app.Post("/products", routes.CreateProduct)
	app.Get("/product/:id", routes.GetProduct)

	app.Listen(":8080")
}
