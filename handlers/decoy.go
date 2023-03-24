package handlers

import (
	"projetoRaul/database"
	"projetoRaul/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// Get Decoy
// @Summary Get  list of Decoy
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string][]string
// @Failure 400
// @Router /decoys [get]
func ListDecoys(c *fiber.Ctx) error {
	decoy := []models.Decoy{}
	database.DB.Db.Find(&decoy)

	return c.Status(200).JSON(decoy)
}

// GetDecoy
// @Summary Insert new Decoy
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string][]string
// @Failure 400
// @Router /decoy [get]
func CreateDecoy(c *fiber.Ctx) error {
	decoy := new(models.Decoy)
	if err := c.BodyParser(decoy); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&decoy)

	return c.Status(200).JSON(decoy)
}

// GetDecoy
// @Summary update a Decoy
// @Description update a Decoy by its ID
// @ID get-Decoy-by-int
// @Accept  json
// @Produce  json
// @Tags Item
// @Param id path int true "Decoy ID"
// @Success 200 {object} map[string][]string
// @Failure 400
// @Router /updateDecoy/{id} [put]
func UpdateDecoy(c *fiber.Ctx) error {
	decoyID := c.Params("id")
	id, err := strconv.Atoi(decoyID)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse id",
		})
	}

	var payload *models.Decoy
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	var decoy models.Decoy
	database.DB.Db.Find(&decoyID, id)
	if int(decoy.ID) != id {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "building not found",
		})
	}

	UpdateDecoy := models.Decoy{
		Name: payload.Name,
	}

	isUpdated := database.DB.Db.Updates(UpdateDecoy)
	if isUpdated == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	database.DB.Db.Where("id= ?", id).Updates(&UpdateDecoy)
	return c.Status(fiber.StatusOK).JSON(UpdateDecoy)
}

// Delete Decoy
// @Summary delete a Decoy
// @Description delete a Decoy by its ID
// @ID delete-decoy-by-int
// @Accept  json
// @Produce  json
// @Tags Decoy
// @Param id path int true "Decoy ID"
// @Success 200 {object} map[string][]string
// @Router /deleteDecoy/{id} [delete]
func DeleteDecoy(c *fiber.Ctx) error {

	decoyID := c.Params("id")

	var decoy *models.Decoy
	result := database.DB.Db.Unscoped().Delete(&decoy, "id= ?", decoyID)

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No decoy with that Id exists"})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error})
	}

	return c.SendStatus(fiber.StatusNoContent)

}
