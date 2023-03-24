package handlers

import (
	"projetoRaul/database"
	"projetoRaul/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// Get Climates
// @Summary Get  list of Climates
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string][]string
// @Failure 400
// @Router /climates [get]
func ListClimate(c *fiber.Ctx) error {
	climate := []models.Climate{}
	database.DB.Db.Find(&climate)

	return c.Status(200).JSON(climate)
}

// GetClimate
// @Summary Insert new Climate
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string][]string
// @Failure 400
// @Router /climate [get]
func CreateClimate(c *fiber.Ctx) error {
	climate := new(models.Climate)
	if err := c.BodyParser(climate); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&climate)

	return c.Status(200).JSON(climate)
}

// GetClimate
// @Summary update a Climate
// @Description update a Climate by its ID
// @ID get-climate-by-int
// @Accept  json
// @Produce  json
// @Tags Item
// @Param id path int true "Climate ID"
// @Success 200 {object} map[string][]string
// @Failure 400
// @Router /updateClimate/{id} [put]
func UpdateClimate(c *fiber.Ctx) error {
	climateID := c.Params("id")
	id, err := strconv.Atoi(climateID)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse id",
		})
	}

	var payload *models.Climate
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	var climate models.Climate
	database.DB.Db.Find(&climateID, id)
	if int(climate.ID) != id {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "building not found",
		})
	}

	UpdateClimate := models.Climate{
		Name: payload.Name,
	}

	isUpdated := database.DB.Db.Updates(UpdateClimate)
	if isUpdated == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	database.DB.Db.Where("id= ?", id).Updates(&UpdateClimate)
	return c.Status(fiber.StatusOK).JSON(UpdateClimate)
}

// Delete Climate
// @Summary delete a Climate
// @Description delete a Climate by its ID
// @ID delete-Climate-by-int
// @Accept  json
// @Produce  json
// @Tags Item
// @Param id path int true "Climate ID"
// @Success 200 {object} map[string][]string
// @Router /deleteClimate/{id} [delete]
func DeleteClimate(c *fiber.Ctx) error {

	climateID := c.Params("id")

	var climate *models.Climate
	result := database.DB.Db.Unscoped().Delete(&climate, "id= ?", climateID)

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No asset with that Id exists"})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error})
	}

	return c.SendStatus(fiber.StatusNoContent)

}
