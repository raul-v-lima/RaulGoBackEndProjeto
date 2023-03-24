package handlers

import (
	"projetoRaul/database"
	"projetoRaul/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// Get Terrain
// @Summary Get  list of Terrain
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string][]string
// @Failure 400
// @Router /terrains [get]
func ListTerrains(c *fiber.Ctx) error {
	terrain := []models.Terrain{}
	database.DB.Db.Find(&terrain)

	return c.Status(200).JSON(terrain)
}

// GetTerrain
// @Summary Insert new Terrain
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string][]string
// @Failure 400
// @Router /terrain [get]
func CreateTerrain(c *fiber.Ctx) error {
	terrain := new(models.Terrain)
	if err := c.BodyParser(terrain); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&terrain)

	return c.Status(200).JSON(terrain)
}

// GetTerrain
// @Summary update a Terrain
// @Description update a Terrain by its ID
// @ID get-Terrain-by-int
// @Accept  json
// @Produce  json
// @Tags Terrain
// @Param id path int true "Terrain ID"
// @Success 200 {object} map[string][]string
// @Failure 400
// @Router /updateTerrain/{id} [put]
func UpdateTerrain(c *fiber.Ctx) error {
	terrainID := c.Params("id")
	id, err := strconv.Atoi(terrainID)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse id",
		})
	}

	var payload *models.Terrain
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	var terrain models.Terrain
	database.DB.Db.Find(&terrainID, id)
	if int(terrain.ID) != id {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "building not found",
		})
	}

	UpdateTerrain := models.Terrain{
		Name: payload.Name,
	}

	isUpdated := database.DB.Db.Updates(UpdateTerrain)
	if isUpdated == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	database.DB.Db.Where("id= ?", id).Updates(&UpdateTerrain)
	return c.Status(fiber.StatusOK).JSON(UpdateTerrain)
}

// Delete Terrain
// @Summary delete a Terrain
// @Description delete a Terrain by its ID
// @ID delete-terrain-by-int
// @Accept  json
// @Produce  json
// @Tags Terrain
// @Param id path int true "Terrain ID"
// @Success 200 {object} map[string][]string
// @Router /deleteTerrain/{id} [delete]
func DeleteTerrain(c *fiber.Ctx) error {

	terrainID := c.Params("id")

	var terrain *models.Terrain
	result := database.DB.Db.Unscoped().Delete(&terrain, "id= ?", terrainID)

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No terrain with that Id exists"})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error})
	}

	return c.SendStatus(fiber.StatusNoContent)

}
