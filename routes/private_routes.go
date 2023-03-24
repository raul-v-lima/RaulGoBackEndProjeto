package routes

import (
	"projetoRaul/handlers"
	"projetoRaul/middleware"

	"github.com/gofiber/fiber/v2"
)

func PrivateRoutes(a *fiber.App) {
	//create routes group

	route := a.Group("")

	route.Post("/asset/", middleware.JWTProtected(), handlers.CreateAsset)
	route.Post("/building", middleware.JWTProtected(), handlers.CreateBuilding)
	route.Post("/character", middleware.JWTProtected(), handlers.CreateCaracter)
	route.Post("/climate", middleware.JWTProtected(), handlers.CreateClimate)
	route.Post("/creature", middleware.JWTProtected(), handlers.CreateCreature)
	route.Post("/decoy", middleware.JWTProtected(), handlers.CreateDecoy)
	route.Post("/endemicCreature", middleware.JWTProtected(), handlers.CreateEndemicCreature)
	route.Post("/endemicVegetation", middleware.JWTProtected(), handlers.CreateEndemicVegetation)
	route.Post("/environment", middleware.JWTProtected(), handlers.CreateEnvironment)
	route.Post("/item", middleware.JWTProtected(), handlers.CreateItem)
	route.Post("/landscape", middleware.JWTProtected(), handlers.CreateLandscape)
	route.Post("/player", middleware.JWTProtected(), handlers.CreatePlayer)
	route.Post("/terrain", middleware.JWTProtected(), handlers.CreateTerrain)
	route.Post("/vegetation", middleware.JWTProtected(), handlers.CreateVegetation)
	route.Post("/waterbody", middleware.JWTProtected(), handlers.CreateWaterBody)

	route.Put("/updateAsset/:id", middleware.JWTProtected(), handlers.UpdateAsset)
	route.Put("/updateBuilding/:id", middleware.JWTProtected(), handlers.UpdateBuilding)
	route.Put("/updateCharacter/:id", middleware.JWTProtected(), handlers.UpdateCharacter)
	route.Put("/updateClimate/:id", middleware.JWTProtected(), handlers.UpdateClimate)
	route.Put("/updateCreature/:id", middleware.JWTProtected(), handlers.UpdateCreature)
	route.Put("/updateDecoy/:id", middleware.JWTProtected(), handlers.UpdateDecoy)
	route.Put("/updateEndemicCreature/:id", middleware.JWTProtected(), handlers.UpdateEndemicCreature)
	route.Put("/updateEndemicVegetation/:id", middleware.JWTProtected(), handlers.UpdateEndemicVegetation)
	route.Put("/updateEnvironment/:id", middleware.JWTProtected(), handlers.UpdateEnvironment)
	route.Put("/updateItem/:id", middleware.JWTProtected(), handlers.UpdateItem)
	route.Put("/updateLandscape/:id", middleware.JWTProtected(), handlers.UpdateLandscape)
	route.Put("/updatePlayer/:id", middleware.JWTProtected(), handlers.UpdatePlayer)
	route.Put("/updateTerrain/:id", middleware.JWTProtected(), handlers.UpdateTerrain)
	route.Put("/updateVegetation/:id", middleware.JWTProtected(), handlers.UpdateVegetation)
	route.Put("/updateWaterbody/:id", middleware.JWTProtected(), handlers.UpdateWaterBody)

	route.Delete("/deleteAsset/:id", middleware.JWTProtected(), handlers.DeleteAsset)
	route.Delete("/deleteBuilding/:id", middleware.JWTProtected(), handlers.DeleteBuilding)
	route.Delete("/deleteCharacter/:id", middleware.JWTProtected(), handlers.DeleteCharacter)
	route.Delete("/deleteClimate/:id", middleware.JWTProtected(), handlers.DeleteClimate)
	route.Delete("/deleteCreature/:id", middleware.JWTProtected(), handlers.DeleteCreature)
	route.Delete("/deleteDecoy/:id", middleware.JWTProtected(), handlers.DeleteDecoy)
	route.Delete("/deleteEndemicCreature/:id", middleware.JWTProtected(), handlers.DeleteEndemicCreature)
	route.Delete("/deleteEndemicVegetation/:id", middleware.JWTProtected(), handlers.DeleteEndemicVegetation)
	route.Delete("/deleteEnvironment/:id", middleware.JWTProtected(), handlers.DeleteEnvironment)
	route.Delete("/deleteItem/:id", middleware.JWTProtected(), handlers.DeleteItem)
	route.Delete("/deleteLandscape/:id", middleware.JWTProtected(), handlers.DeleteLandscape)
	route.Delete("/deletePlayer/:id", middleware.JWTProtected(), handlers.DeletePlayer)
	route.Delete("/deleteTerrain/:id", middleware.JWTProtected(), handlers.DeleteTerrain)
	route.Delete("/deleteVegetation/:id", middleware.JWTProtected(), handlers.DeleteVegetation)
	route.Delete("/deleteWaterbody/:id", middleware.JWTProtected(), handlers.DeleteWaterBody)

}
