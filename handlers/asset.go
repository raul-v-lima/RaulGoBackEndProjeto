package handlers

import (
	"projetoRaul/database"
	"projetoRaul/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// Get Assets
// @Summary Get  list of Assets
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string][]string
// @Failure 400
// @Router /assets [get]
func ListAssets(c *fiber.Ctx) error {
	asset := []models.Asset{}
	database.DB.Db.Find(&asset)

	return c.Status(200).JSON(asset)
}

// @Summary Insert new Asset
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string][]string
// @Failure 400
// @Router /asset [get]
func CreateAsset(c *fiber.Ctx) error {
	asset := new(models.Asset)
	if err := c.BodyParser(asset); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&asset)

	return c.Status(200).JSON(asset)
}

// GetAsset
// @Summary update a Asset
// @Description update a Asset by its ID
// @ID get-Asset-by-int
// @Accept  json
// @Produce  json
// @Tags Asset
// @Param id path int true "Asset ID"
// @Success 200 {object} map[string][]string
// @Failure 400
// @Router /updateAsset/{id} [put]
func UpdateAsset(c *fiber.Ctx) error {
	assetID := c.Params("id")
	id, err := strconv.Atoi(assetID)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse id",
		})
	}

	var payload *models.Asset
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	var asset models.Asset
	database.DB.Db.Find(&asset, id)
	if int(asset.ID) != id {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Asset not found",
		})
	}

	UpdateAsset := models.Asset{
		Name: payload.Name,
	}

	isUpdated := database.DB.Db.Updates(UpdateAsset)
	if isUpdated == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	database.DB.Db.Where("id= ?", id).Updates(&UpdateAsset)
	return c.Status(fiber.StatusOK).JSON(UpdateAsset)
}

// Delete Asset
// @Summary delete a Asset
// @Description delete a Asset by its ID
// @ID delete-Asset-by-int
// @Accept  json
// @Produce  json
// @Tags Asset
// @Param id path int true "Asset ID"
// @Success 200 {object} map[string][]string
// @Router /deleteAsset/{id} [delete]
func DeleteAsset(c *fiber.Ctx) error {

	assetID := c.Params("id")

	var asset *models.Asset
	result := database.DB.Db.Unscoped().Delete(&asset, "id= ?", assetID)

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No asset with that Id exists"})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error})
	}

	return c.SendStatus(fiber.StatusNoContent)

}
