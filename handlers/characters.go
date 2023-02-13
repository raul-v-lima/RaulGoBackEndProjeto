package handlers

import (
	"projetoRaul/database"
	"projetoRaul/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// Get Characters
// @Summary Get  list of Characters
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string][]string
// @Failure 400
// @Router / [get]
func ListCharacters(c *fiber.Ctx) error {
	character := []models.Character{}
	database.DB.Db.Find(&character)

	return c.Status(200).JSON(character)
}

// GetCharacters
// @Summary Insert new Character
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string][]string
// @Failure 400
// @Router /character [get]
func CreateCaracters(c *fiber.Ctx) error {
	fact := new(models.Character)
	if err := c.BodyParser(fact); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&fact)

	return c.Status(200).JSON(fact)
}

// GetItem
// @Summary update a Character
// @Description update a Character by its ID
// @ID get-Character-by-int
// @Accept  json
// @Produce  json
// @Tags Item
// @Param id path int true "Character ID"
// @Success 200 {object} map[string][]string
// @Failure 400
// @Router /updateCharacter/{id} [put]
func UpdateCharacter(c *fiber.Ctx) error {
	characterID := c.Params("id")
	id, err := strconv.Atoi(characterID)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse id",
		})
	}

	var payload *models.Character
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	var character models.Character
	database.DB.Db.Find(&character, id)
	if int(character.ID) != id {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Character not found",
		})
	}

	updatecharacter := models.Character{
		Name:      payload.Name,
		Vigor:     payload.Vigor,
		Dexterity: payload.Dexterity,
		Xp:        payload.Xp,
		Level:     payload.Level,
		Mana:      payload.Mana,
	}

	isUpdated := database.DB.Db.Updates(updatecharacter)
	if isUpdated == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	database.DB.Db.Where("id= ?", id).Updates(&updatecharacter)
	return c.Status(fiber.StatusOK).JSON(updatecharacter)
}

// Delete Character
// @Summary delete a Character
// @Description delete a Character by its ID
// @ID delete-Character-by-int
// @Accept  json
// @Produce  json
// @Tags Item
// @Param id path int true "Character ID"
// @Success 200 {object} map[string][]string
// @Router /deleteCharacter/{id} [delete]
func DeleteCharacter(c *fiber.Ctx) error {

	characterID := c.Params("id")

	var character *models.Character
	result := database.DB.Db.Unscoped().Delete(&character, "id= ?", characterID)

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No character with that Id exists"})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error})
	}

	return c.SendStatus(fiber.StatusNoContent)

}
