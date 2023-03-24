package handlers

import (
	"projetoRaul/database"
	"projetoRaul/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// Get Player
// @Summary Get  list of Player
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string][]string
// @Failure 400
// @Router /players [get]
func ListPlayers(c *fiber.Ctx) error {
	player := []models.Player{}
	database.DB.Db.Find(&player)

	return c.Status(200).JSON(player)
}

// GetPlayer
// @Summary Insert new Player
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string][]string
// @Failure 400
// @Router /player [get]
func CreatePlayer(c *fiber.Ctx) error {
	player := new(models.Player)
	if err := c.BodyParser(player); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&player)

	return c.Status(200).JSON(player)
}

// GetPlayer
// @Summary update a Player
// @Description update a Player by its ID
// @ID get-Player-by-int
// @Accept  json
// @Produce  json
// @Tags Player
// @Param id path int true "Player ID"
// @Success 200 {object} map[string][]string
// @Failure 400
// @Router /updatePlayer/{id} [put]
func UpdatePlayer(c *fiber.Ctx) error {
	playerID := c.Params("id")
	id, err := strconv.Atoi(playerID)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse id",
		})
	}

	var payload *models.Player
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	var player models.Player
	database.DB.Db.Find(&playerID, id)
	if int(player.ID) != id {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "building not found",
		})
	}

	UpdatePlayer := models.Player{
		Name: payload.Name,
	}

	isUpdated := database.DB.Db.Updates(UpdatePlayer)
	if isUpdated == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	database.DB.Db.Where("id= ?", id).Updates(&UpdatePlayer)
	return c.Status(fiber.StatusOK).JSON(UpdatePlayer)
}

// Delete Player
// @Summary delete a Player
// @Description delete a Player by its ID
// @ID delete-reature-by-int
// @Accept  json
// @Produce  json
// @Tags Player
// @Param id path int true "Player ID"
// @Success 200 {object} map[string][]string
// @Router /deletePlayer/{id} [delete]
func DeletePlayer(c *fiber.Ctx) error {

	playerID := c.Params("id")

	var player *models.Player
	result := database.DB.Db.Unscoped().Delete(&player, "id= ?", playerID)

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No player with that Id exists"})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error})
	}

	return c.SendStatus(fiber.StatusNoContent)

}
