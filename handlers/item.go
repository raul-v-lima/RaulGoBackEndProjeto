package handlers

import (
	"projetoRaul/database"
	"projetoRaul/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// Get Item
// @Summary Get  list of Item
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string][]string
// @Failure 400
// @Router /items [get]
func ListItems(c *fiber.Ctx) error {
	item := []models.Item{}
	database.DB.Db.Find(&item)

	return c.Status(200).JSON(item)
}

// GetItem
// @Summary Insert new Item
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string][]string
// @Failure 400
// @Router /item [get]
func CreateItem(c *fiber.Ctx) error {
	item := new(models.Item)
	if err := c.BodyParser(item); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&item)

	return c.Status(200).JSON(item)
}

// GetItem
// @Summary update a Item
// @Description update a Item by its ID
// @ID get-Item-by-int
// @Accept  json
// @Produce  json
// @Tags Item
// @Param id path int true "Item ID"
// @Success 200 {object} map[string][]string
// @Failure 400
// @Router /updateItem/{id} [put]
func UpdateItem(c *fiber.Ctx) error {
	itemID := c.Params("id")
	id, err := strconv.Atoi(itemID)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse id",
		})
	}

	var payload *models.Item
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	var item models.Item
	database.DB.Db.Find(&itemID, id)
	if int(item.ID) != id {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "building not found",
		})
	}

	UpdateItem := models.Item{
		Name: payload.Name,
	}

	isUpdated := database.DB.Db.Updates(UpdateItem)
	if isUpdated == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	database.DB.Db.Where("id= ?", id).Updates(&UpdateItem)
	return c.Status(fiber.StatusOK).JSON(UpdateItem)
}

// Delete Item
// @Summary delete a Item
// @Description delete a Item by its ID
// @ID delete-item-by-int
// @Accept  json
// @Produce  json
// @Tags Item
// @Param id path int true "Item ID"
// @Success 200 {object} map[string][]string
// @Router /deleteItem/{id} [delete]
func DeleteItem(c *fiber.Ctx) error {

	itemID := c.Params("id")

	var item *models.Item
	result := database.DB.Db.Unscoped().Delete(&item, "id= ?", itemID)

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No item with that Id exists"})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error})
	}

	return c.SendStatus(fiber.StatusNoContent)

}
