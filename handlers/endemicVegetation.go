package handlers

import (
	"projetoRaul/database"
	"projetoRaul/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// Get EndemicVegetation
// @Summary Get  list of EndemicVegetation
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string][]string
// @Failure 400
// @Router /endemicVegetations [get]
func ListEndemicVegetations(c *fiber.Ctx) error {
	endemicVegetation := []models.EndemicVegetation{}
	database.DB.Db.Find(&endemicVegetation)

	return c.Status(200).JSON(endemicVegetation)
}

// GetEndemicVegetation
// @Summary Insert new EndemicVegetation
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string][]string
// @Failure 400
// @Router /endemicVegetation [get]
func CreateEndemicVegetation(c *fiber.Ctx) error {
	endemicVegetation := new(models.EndemicVegetation)
	if err := c.BodyParser(endemicVegetation); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&endemicVegetation)

	return c.Status(200).JSON(endemicVegetation)
}

// GetEndemicVegetation
// @Summary update a EndemicVegetation
// @Description update a EndemicVegetation by its ID
// @ID get-EndemicVegetation-by-int
// @Accept  json
// @Produce  json
// @Tags Item
// @Param id path int true "EndemicVegetation ID"
// @Success 200 {object} map[string][]string
// @Failure 400
// @Router /updateEndemicVegetation/{id} [put]
func UpdateEndemicVegetation(c *fiber.Ctx) error {
	endemicVegetationID := c.Params("id")
	id, err := strconv.Atoi(endemicVegetationID)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse id",
		})
	}

	var payload *models.EndemicVegetation
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	var endemicVegetation models.EndemicVegetation
	database.DB.Db.Find(&endemicVegetationID, id)
	if int(endemicVegetation.ID) != id {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "building not found",
		})
	}

	UpdateEndemicVegetation := models.EndemicVegetation{
		Name: payload.Name,
	}

	isUpdated := database.DB.Db.Updates(UpdateEndemicVegetation)
	if isUpdated == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	database.DB.Db.Where("id= ?", id).Updates(&UpdateEndemicVegetation)
	return c.Status(fiber.StatusOK).JSON(UpdateEndemicVegetation)
}

// Delete EndemicVegetation
// @Summary delete a EndemicVegetation
// @Description delete a EndemicVegetation by its ID
// @ID delete-endemicVegetations-by-int
// @Accept  json
// @Produce  json
// @Tags EndemicVegetation
// @Param id path int true "EndemicVegetation ID"
// @Success 200 {object} map[string][]string
// @Router /deleteEndemicVegetation/{id} [delete]
func DeleteEndemicVegetation(c *fiber.Ctx) error {

	endemicVegetationID := c.Params("id")

	var endemicVegetation *models.EndemicVegetation
	result := database.DB.Db.Unscoped().Delete(&endemicVegetation, "id= ?", endemicVegetationID)

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No endemicVegetation with that Id exists"})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error})
	}

	return c.SendStatus(fiber.StatusNoContent)

}
