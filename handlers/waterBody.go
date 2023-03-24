package handlers

import (
	"projetoRaul/database"
	"projetoRaul/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// Get WaterBody
// @Summary Get  list of WaterBody
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string][]string
// @Failure 400
// @Router /waterbodies [get]
func ListWaterBodies(c *fiber.Ctx) error {
	waterBody := []models.WaterBody{}
	database.DB.Db.Find(&waterBody)

	return c.Status(200).JSON(waterBody)
}

// GetWaterBody
// @Summary Insert new WaterBody
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string][]string
// @Failure 400
// @Router /waterBody [get]
func CreateWaterBody(c *fiber.Ctx) error {
	waterBody := new(models.WaterBody)
	if err := c.BodyParser(waterBody); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&waterBody)

	return c.Status(200).JSON(waterBody)
}

// GetWaterBody
// @Summary update a WaterBody
// @Description update a WaterBody by its ID
// @ID get-WaterBody-by-int
// @Accept  json
// @Produce  json
// @Tags WaterBody
// @Param id path int true "WaterBody ID"
// @Success 200 {object} map[string][]string
// @Failure 400
// @Router /updateWaterBody/{id} [put]
func UpdateWaterBody(c *fiber.Ctx) error {
	waterBodyID := c.Params("id")
	id, err := strconv.Atoi(waterBodyID)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse id",
		})
	}

	var payload *models.WaterBody
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	var waterBody models.WaterBody
	database.DB.Db.Find(&waterBodyID, id)
	if int(waterBody.ID) != id {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "building not found",
		})
	}

	UpdateWaterBody := models.WaterBody{
		Name: payload.Name,
	}

	isUpdated := database.DB.Db.Updates(UpdateWaterBody)
	if isUpdated == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	database.DB.Db.Where("id= ?", id).Updates(&UpdateWaterBody)
	return c.Status(fiber.StatusOK).JSON(UpdateWaterBody)
}

// Delete WaterBody
// @Summary delete a WaterBody
// @Description delete a WaterBody by its ID
// @ID delete-waterBody-by-int
// @Accept  json
// @Produce  json
// @Tags WaterBody
// @Param id path int true "WaterBody ID"
// @Success 200 {object} map[string][]string
// @Router /deleteWaterBody/{id} [delete]
func DeleteWaterBody(c *fiber.Ctx) error {

	waterBodyID := c.Params("id")

	var waterBody *models.WaterBody
	result := database.DB.Db.Unscoped().Delete(&waterBody, "id= ?", waterBodyID)

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No waterBody with that Id exists"})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error})
	}

	return c.SendStatus(fiber.StatusNoContent)

}
