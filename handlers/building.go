package handlers

import (
	"projetoRaul/database"
	"projetoRaul/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// Get Building
// @Summary Get  list of Building
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string][]string
// @Failure 400
// @Router /buildings [get]
func ListBuildings(c *fiber.Ctx) error {
	building := []models.Building{}
	database.DB.Db.Find(&building)

	return c.Status(200).JSON(building)
}

// GetBuilding
// @Summary Insert new Building
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string][]string
// @Failure 400
// @Router /building [get]
func CreateBuilding(c *fiber.Ctx) error {
	building := new(models.Building)
	if err := c.BodyParser(building); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&building)

	return c.Status(200).JSON(building)
}

// GetBuilding
// @Summary update a Building
// @Description update a Building by its ID
// @ID get-building-by-int
// @Accept  json
// @Produce  json
// @Tags Building
// @Param id path int true "Building ID"
// @Success 200 {object} map[string][]string
// @Failure 400
// @Router /updateBuilding/{id} [put]
func UpdateBuilding(c *fiber.Ctx) error {
	buildingID := c.Params("id")
	id, err := strconv.Atoi(buildingID)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse id",
		})
	}

	var payload *models.Building
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	var building models.Building
	database.DB.Db.Find(&buildingID, id)
	if int(building.ID) != id {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "building not found",
		})
	}

	UpdateBuilding := models.Building{
		Name: payload.Name,
	}

	isUpdated := database.DB.Db.Updates(UpdateBuilding)
	if isUpdated == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	database.DB.Db.Where("id= ?", id).Updates(&UpdateBuilding)
	return c.Status(fiber.StatusOK).JSON(UpdateBuilding)
}

// Delete Building
// @Summary delete a Building
// @Description delete a Building by its ID
// @ID delete-building-by-int
// @Accept  json
// @Produce  json
// @Tags Item
// @Param id path int true "Building ID"
// @Success 200 {object} map[string][]string
// @Router /deleteBuilding/{id} [delete]
func DeleteBuilding(c *fiber.Ctx) error {

	buildingID := c.Params("id")

	var building *models.Building
	result := database.DB.Db.Unscoped().Delete(&building, "id= ?", buildingID)

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No building with that Id exists"})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error})
	}

	return c.SendStatus(fiber.StatusNoContent)

}
