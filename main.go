package main

import (
	"projetoRaul/database"
	"projetoRaul/handlers"

	"github.com/gofiber/fiber/v2"

	_ "projetoRaul/docs"

	swagger "github.com/arsmn/fiber-swagger/v2"
)

// @title Fiber Example API
// @version 1.0
// @description  swagger Projeto Raul
// @contact.name API Support
// @host localhost:5000
// @BasePath /

func main() {

	database.ConnectDb()
	// Create new Fiber application
	app := fiber.New()

	app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL:         "/swagger/doc.json",
		DeepLinking: false,
	}))

	app.Get("/", handlers.ListCharacters)
	app.Post("/character", handlers.CreateCaracters)
	app.Put("/updateCharacter/:id", handlers.UpdateCharacter)
	app.Delete("/deleteCharacter/:id", handlers.DeleteCharacter)

	app.Listen(":5000")

	//implement oauth
}
