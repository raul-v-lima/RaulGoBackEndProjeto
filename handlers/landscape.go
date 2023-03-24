package handlers

import (
	"projetoRaul/database"
	"projetoRaul/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// Get Landscape
// @Summary Get  list of Landscape
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string][]string
// @Failure 400
// @Router /landscapes [get]
func ListLandscapes(c *fiber.Ctx) error {
	landscape := []models.Landscape{}
	database.DB.Db.Find(&landscape)

	return c.Status(200).JSON(landscape)
}

// GetLandscape
// @Summary Insert new Landscape
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string][]string
// @Failure 400
// @Router /landscape [get]
func CreateLandscape(c *fiber.Ctx) error {
	landscape := new(models.Landscape)
	if err := c.BodyParser(landscape); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&landscape)

	return c.Status(200).JSON(landscape)
}

// GetLandscape
// @Summary update a Landscape
// @Description update a Landscape by its ID
// @ID get-Landscape-by-int
// @Accept  json
// @Produce  json
// @Tags Landscape
// @Param id path int true "Landscape ID"
// @Success 200 {object} map[string][]string
// @Failure 400
// @Router /updateLandscape/{id} [put]
func UpdateLandscape(c *fiber.Ctx) error {
	landscapeID := c.Params("id")
	id, err := strconv.Atoi(landscapeID)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse id",
		})
	}

	var payload *models.Landscape
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	var landscape models.Landscape
	database.DB.Db.Find(&landscapeID, id)
	if int(landscape.ID) != id {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "building not found",
		})
	}

	UpdateLandscape := models.Landscape{
		Name: payload.Name,
	}

	isUpdated := database.DB.Db.Updates(UpdateLandscape)
	if isUpdated == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	database.DB.Db.Where("id= ?", id).Updates(&UpdateLandscape)
	return c.Status(fiber.StatusOK).JSON(UpdateLandscape)
}

// Delete Landscape
// @Summary delete a Landscape
// @Description delete a Landscape by its ID
// @ID delete-landscape-by-int
// @Accept  json
// @Produce  json
// @Tags Landscape
// @Param id path int true "Landscape ID"
// @Success 200 {object} map[string][]string
// @Router /deleteLandscape/{id} [delete]
func DeleteLandscape(c *fiber.Ctx) error {

	landscapeID := c.Params("id")

	var landscape *models.Landscape
	result := database.DB.Db.Unscoped().Delete(&landscape, "id= ?", landscapeID)

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No landscape with that Id exists"})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error})
	}

	return c.SendStatus(fiber.StatusNoContent)

}
