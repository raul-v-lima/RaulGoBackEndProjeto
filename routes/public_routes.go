package routes

import (
	"projetoRaul/handlers"

	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(a *fiber.App) {

	route := a.Group("")

	route.Get("/assets", handlers.ListAssets)
	route.Get("/buildings", handlers.ListBuildings)
	route.Get("/characters", handlers.ListCharacters)
	route.Get("/climates", handlers.ListClimate)
	route.Get("/creatures", handlers.ListCreatures)
	route.Get("/decoys", handlers.ListDecoys)
	route.Get("/endemicCreatures", handlers.ListEndemicCreatures)
	route.Get("/endemicVegetations", handlers.ListEndemicVegetations)
	route.Get("/environments", handlers.ListEnvironments)
	route.Get("/items", handlers.ListItems)
	route.Get("/landscapes", handlers.ListLandscapes)
	route.Get("/players", handlers.ListPlayers)
	route.Get("/terrains", handlers.ListTerrains)
	route.Get("/vegetation", handlers.ListVegetations)
	route.Get("/waterbodies", handlers.ListWaterBodies)
	route.Get("/token/new", handlers.GetNewAccessToken)
}
