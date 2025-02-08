package controllers

import (
	db "Exercise/OrderAPI/config"
	"Exercise/OrderAPI/models"

	"github.com/gofiber/fiber/v2"
)

func PlaceOrder(c *fiber.Ctx) error {
	var order models.Order
	if err := c.BodyParser(&order); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}

	if order.Items == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "No items selected!",
		})
	}

	order.Status = "Pending"
	db.DB.Create(&order)
	return c.Status(201).JSON(order)
}

func ApproveOrder(c *fiber.Ctx) error {
	CustomerId := c.Params("CustomerId")
	var order models.Order

	if err := db.DB.First(&order, CustomerId).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Order not found",
		})
	}

	order.Status = "Approved"
	db.DB.Save(&order)
	return c.JSON(order)

}

func CancelOrder(c * fiber.Ctx) error{
	CustomerId := c.Params("CustomerId")
	var order models.Order

	if err := db.DB.First(&order,CustomerId).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error":"Order not found",
		})
	}

	order.Status = "Cancel"
	db.DB.Save(&order)
	return c.JSON(order)
}

func GetBill(c *fiber.Ctx) error {
	CustomerId := c.Params("Customerid")
	var order models.Order

	if err := db.DB.First(&order, CustomerId).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Order not found",
		})
	}

	if order.Status != "Approved" {
		return c.Status(400).JSON(fiber.Map{
			"error": "Order is not approved!",
		})
	}
	return c.JSON(fiber.Map{

		"CustomerId": order.CustomerId,
		"totalprice": order.TotalPrice,
		"items":      order.Items,
	})
}
