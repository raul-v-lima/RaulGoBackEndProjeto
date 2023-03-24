package handlers

import (
	"projetoRaul/database"
	"projetoRaul/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// Get EndemicCreature
// @Summary Get  list of EndemicCreature
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string][]string
// @Failure 400
// @Router /endemicCreatures [get]
func ListEndemicCreatures(c *fiber.Ctx) error {
	endemicCreature := []models.EndemicCreature{}
	database.DB.Db.Find(&endemicCreature)

	return c.Status(200).JSON(endemicCreature)
}

// GetEndemicCreature
// @Summary Insert new EndemicCreature
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string][]string
// @Failure 400
// @Router /endemicCreature [get]
func CreateEndemicCreature(c *fiber.Ctx) error {
	endemicCreature := new(models.EndemicCreature)
	if err := c.BodyParser(endemicCreature); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&endemicCreature)

	return c.Status(200).JSON(endemicCreature)
}

// GetEndemicCreature
// @Summary update a EndemicCreature
// @Description update a EndemicCreature by its ID
// @ID get-EndemicCreature-by-int
// @Accept  json
// @Produce  json
// @Tags Item
// @Param id path int true "EndemicCreature ID"
// @Success 200 {object} map[string][]string
// @Failure 400
// @Router /updateEndemicCreature/{id} [put]
func UpdateEndemicCreature(c *fiber.Ctx) error {
	endemicCreatureID := c.Params("id")
	id, err := strconv.Atoi(endemicCreatureID)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse id",
		})
	}

	var payload *models.EndemicCreature
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	var endemicCreature models.EndemicCreature
	database.DB.Db.Find(&endemicCreatureID, id)
	if int(endemicCreature.ID) != id {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "building not found",
		})
	}

	UpdateEndemicCreature := models.EndemicCreature{
		Name: payload.Name,
	}

	isUpdated := database.DB.Db.Updates(UpdateEndemicCreature)
	if isUpdated == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	database.DB.Db.Where("id= ?", id).Updates(&UpdateEndemicCreature)
	return c.Status(fiber.StatusOK).JSON(UpdateEndemicCreature)
}

// Delete EndemicCreature
// @Summary delete a EndemicCreature
// @Description delete a EndemicCreature by its ID
// @ID delete-endemicCreature-by-int
// @Accept  json
// @Produce  json
// @Tags EndemicCreature
// @Param id path int true "EndemicCreature ID"
// @Success 200 {object} map[string][]string
// @Router /deleteEndemicCreature/{id} [delete]
func DeleteEndemicCreature(c *fiber.Ctx) error {

	endemicCreatureID := c.Params("id")

	var endemicCreature *models.EndemicCreature
	result := database.DB.Db.Unscoped().Delete(&endemicCreature, "id= ?", endemicCreatureID)

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No endemicCreature with that Id exists"})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error})
	}

	return c.SendStatus(fiber.StatusNoContent)

}
