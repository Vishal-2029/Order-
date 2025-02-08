package controllers

import (
	db "Exercise/OrderAPI/config"
	"Exercise/OrderAPI/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func CreateItems(c *fiber.Ctx) error {
	var Item models.Items

	if err := c.BodyParser(&Item); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse JSON",
		})
	}
	items := models.Items{
		ItemName: Item.ItemName,
		Price:    Item.Price,
		Stock:    Item.Stock,
	}

	db.DB.Create(&items)
	return c.Status(200).JSON(fiber.Map{
		"Success": true,
		"Message": items.ItemName + " is Added",
		"Data":    items,
	})
}

type items struct {
	ItemName string  `json:"customerName"`
	Price    float64 `json:"price"`
	Stock    int     `json:"stock"`
}

func ItemList(c *fiber.Ctx) error {
	limit, _ := strconv.Atoi(c.Query("limit"))
	skip, _ := strconv.Atoi(c.Query("skip"))
	var count int64
	var Item []items
	db.DB.Select("*").Limit(limit).Offset(skip).Find(&Item).Count(&count)
	metaMap := map[string]interface{}{
		"total": count,
		"limit": limit,
		"skip":  skip,
	}
	ItemData := map[string]interface{}{
		"Item": Item,
		"meta": metaMap,
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"Message": "Add Item Successfully",
		"data":    ItemData,
	})
}

func UpdateItems(c *fiber.Ctx) error {
	ItemId := c.Params("ItemId")
	var Item models.Items

	if err := db.DB.First(&Item, ItemId).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Item is Not Found",
		})
	}

	var updateItemId models.Items
	if err := c.BodyParser(&updateItemId); err != nil || updateItemId.ItemName == "" {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Valid ID name is required",
		})
	}

	Item.ItemName = updateItemId.ItemName
	db.DB.Save(&Item)

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Update ID successfully",
		"data":    Item,
	})
}

func DeleteItem(c *fiber.Ctx) error {
	ItemId := c.Params("ItemID")
	var Item models.Items

	if err := db.DB.First(&Item, ItemId).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "ID is Not Found",
		})
	}

	db.DB.Delete(&Item)

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Item is Deleted successfully",
	})
}
