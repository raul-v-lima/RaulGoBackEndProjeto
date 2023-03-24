package handlers

import (
	"projetoRaul/database"
	"projetoRaul/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// Get Vegetation
// @Summary Get  list of Vegetation
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string][]string
// @Failure 400
// @Router /vegetations [get]
func ListVegetations(c *fiber.Ctx) error {
	vegetation := []models.Vegetation{}
	database.DB.Db.Find(&vegetation)

	return c.Status(200).JSON(vegetation)
}

// GetVegetation
// @Summary Insert new Vegetation
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string][]string
// @Failure 400
// @Router /vegetation [get]
func CreateVegetation(c *fiber.Ctx) error {
	vegetation := new(models.Vegetation)
	if err := c.BodyParser(vegetation); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&vegetation)

	return c.Status(200).JSON(vegetation)
}

// GetVegetation
// @Summary update a Vegetation
// @Description update a Vegetation by its ID
// @ID get-Vegetation-by-int
// @Accept  json
// @Produce  json
// @Tags Vegetation
// @Param id path int true "Vegetation ID"
// @Success 200 {object} map[string][]string
// @Failure 400
// @Router /updateVegetation/{id} [put]
func UpdateVegetation(c *fiber.Ctx) error {
	vegetationID := c.Params("id")
	id, err := strconv.Atoi(vegetationID)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse id",
		})
	}

	var payload *models.Vegetation
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	var vegetation models.Vegetation
	database.DB.Db.Find(&vegetationID, id)
	if int(vegetation.ID) != id {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "building not found",
		})
	}

	UpdateVegetation := models.Vegetation{
		Name: payload.Name,
	}

	isUpdated := database.DB.Db.Updates(UpdateVegetation)
	if isUpdated == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	database.DB.Db.Where("id= ?", id).Updates(&UpdateVegetation)
	return c.Status(fiber.StatusOK).JSON(UpdateVegetation)
}

// Delete Vegetation
// @Summary delete a Vegetation
// @Description delete a Vegetation by its ID
// @ID delete-vegetation-by-int
// @Accept  json
// @Produce  json
// @Tags Vegetation
// @Param id path int true "Vegetation ID"
// @Success 200 {object} map[string][]string
// @Router /deleteVegetation/{id} [delete]
func DeleteVegetation(c *fiber.Ctx) error {

	vegetationID := c.Params("id")

	var vegetation *models.Vegetation
	result := database.DB.Db.Unscoped().Delete(&vegetation, "id= ?", vegetationID)

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No vegetation with that Id exists"})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error})
	}

	return c.SendStatus(fiber.StatusNoContent)

}
