package handlers

import (
	"projetoRaul/database"
	"projetoRaul/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// Get Environment
// @Summary Get  list of Environment
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string][]string
// @Failure 400
// @Router /environments [get]
func ListEnvironments(c *fiber.Ctx) error {
	environment := []models.Environment{}
	database.DB.Db.Find(&environment)

	return c.Status(200).JSON(environment)
}

// GetEnvironment
// @Summary Insert new Environment
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string][]string
// @Failure 400
// @Router /environment [get]
func CreateEnvironment(c *fiber.Ctx) error {
	environment := new(models.Environment)
	if err := c.BodyParser(environment); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&environment)

	return c.Status(200).JSON(environment)
}

// GetEnvironment
// @Summary update a Environment
// @Description update a Environment by its ID
// @ID get-Environment-by-int
// @Accept  json
// @Produce  json
// @Tags Item
// @Param id path int true "Environment ID"
// @Success 200 {object} map[string][]string
// @Failure 400
// @Router /updateEnvironment/{id} [put]
func UpdateEnvironment(c *fiber.Ctx) error {
	environmentID := c.Params("id")
	id, err := strconv.Atoi(environmentID)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse id",
		})
	}

	var payload *models.Environment
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	var environment models.Environment
	database.DB.Db.Find(&environmentID, id)
	if int(environment.ID) != id {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "building not found",
		})
	}

	UpdateEnvironment := models.Environment{
		Name: payload.Name,
	}

	isUpdated := database.DB.Db.Updates(UpdateEnvironment)
	if isUpdated == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	database.DB.Db.Where("id= ?", id).Updates(&UpdateEnvironment)
	return c.Status(fiber.StatusOK).JSON(UpdateEnvironment)
}

// Delete Environment
// @Summary delete a Environment
// @Description delete a Environment by its ID
// @ID delete-environment-by-int
// @Accept  json
// @Produce  json
// @Tags Environment
// @Param id path int true "Environment ID"
// @Success 200 {object} map[string][]string
// @Router /deleteEnvironment/{id} [delete]
func DeleteEnvironment(c *fiber.Ctx) error {

	environmentID := c.Params("id")

	var environment *models.Environment
	result := database.DB.Db.Unscoped().Delete(&environment, "id= ?", environmentID)

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No environment with that Id exists"})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error})
	}

	return c.SendStatus(fiber.StatusNoContent)

}
