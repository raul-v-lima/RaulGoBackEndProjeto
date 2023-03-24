package main

import (
	"projetoRaul/configs"
	"projetoRaul/database"
	"projetoRaul/middleware"
	"projetoRaul/routes"
	"projetoRaul/utils"

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

	config := configs.FiberConfig()
	// Create new Fiber application
	app := fiber.New(config)
	middleware.FiberMiddleware(app)

	routes.PublicRoutes(app)  // Register a public routes for app.
	routes.PrivateRoutes(app) // Register a private routes for app.

	app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL:         "/swagger/doc.json",
		DeepLinking: false,
	}))

	//	app.Get("/", handlers.ListCharacters)
	// app.Post("/character", handlers.CreateCaracters)
	//	app.Put("/updateCharacter/:id", handlers.UpdateCharacter)
	//app.Delete("/deleteCharacter/:id", handlers.DeleteCharacter)

	utils.StartServer(app)

	//implement oauth
}
