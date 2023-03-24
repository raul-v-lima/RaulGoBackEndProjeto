package handlers

import (
	"projetoRaul/database"
	"projetoRaul/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// Get Creature
// @Summary Get  list of Creature
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string][]string
// @Failure 400
// @Router /creatures [get]
func ListCreatures(c *fiber.Ctx) error {
	creatures := []models.Creature{}
	database.DB.Db.Find(&creatures)

	return c.Status(200).JSON(creatures)
}

// GetCreature
// @Summary Insert new Creature
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string][]string
// @Failure 400
// @Router /creature [get]
func CreateCreature(c *fiber.Ctx) error {
	creature := new(models.Creature)
	if err := c.BodyParser(creature); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&creature)

	return c.Status(200).JSON(creature)
}

// GetCreature
// @Summary update a Creature
// @Description update a Creature by its ID
// @ID get-Creature-by-int
// @Accept  json
// @Produce  json
// @Tags Item
// @Param id path int true "Creature ID"
// @Success 200 {object} map[string][]string
// @Failure 400
// @Router /updateCreature/{id} [put]
func UpdateCreature(c *fiber.Ctx) error {
	creatureID := c.Params("id")
	id, err := strconv.Atoi(creatureID)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse id",
		})
	}

	var payload *models.Creature
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	var creature models.Creature
	database.DB.Db.Find(&creatureID, id)
	if int(creature.ID) != id {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "building not found",
		})
	}

	UpdateCreature := models.Creature{
		Name: payload.Name,
	}

	isUpdated := database.DB.Db.Updates(UpdateCreature)
	if isUpdated == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	database.DB.Db.Where("id= ?", id).Updates(&UpdateCreature)
	return c.Status(fiber.StatusOK).JSON(UpdateCreature)
}

// Delete Creature
// @Summary delete a Creature
// @Description delete a Creature by its ID
// @ID delete-creature-by-int
// @Accept  json
// @Produce  json
// @Tags Creature
// @Param id path int true "Creature ID"
// @Success 200 {object} map[string][]string
// @Router /deleteCreature/{id} [delete]
func DeleteCreature(c *fiber.Ctx) error {

	creatureID := c.Params("id")

	var creature *models.Creature
	result := database.DB.Db.Unscoped().Delete(&creature, "id= ?", creatureID)

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No asset with that Id exists"})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error})
	}

	return c.SendStatus(fiber.StatusNoContent)

}
