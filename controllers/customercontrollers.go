package controllers

import (
	db "Exercise/OrderAPI/config"
	"Exercise/OrderAPI/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type RequestBody struct {
	CustomerName string `json:"customerName"`
	Passcode     string `json:"passcode"`
}

func CreateCustomers(c *fiber.Ctx) error {
	var body RequestBody

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse JSON",
		})
	}

	Customer := models.Customer{
		CustomerName: body.CustomerName,
		Passcode:     body.Passcode,
	}

	db.DB.Create(&Customer)
	return c.Status(200).JSON(fiber.Map{
		"Success": true,
		"Message": body.CustomerName + " Added",
		"Data":    Customer,
	})
}

type Customer struct {
	CustomerId   uint   `json:"customeroId"`
	CustomerName string `json:"customerName"`
}

func CustomerList(c *fiber.Ctx) error {
	limit, _ := strconv.Atoi(c.Query("limit"))
	skip, _ := strconv.Atoi(c.Query("skip"))
	var count int64
	var customer []Customer
	db.DB.Select("*").Limit(limit).Offset(skip).Find(&customer).Count(&count)
	metaMap := map[string]interface{}{
		"total": count,
		"limit": limit,
		"skip":  skip,
	}
	customerData := map[string]interface{}{
		"customer": customer,
		"meta":     metaMap,
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"Message": "Success",
		"data":    customerData,
	})
}

func UpdateCustomer(c *fiber.Ctx) error {
	CustomerId := c.Params("CustomerId")
	var Customer models.Customer

	if err := db.DB.First(&Customer, CustomerId).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Customer is Not Found",
		})
	}

	var updateCustomerId models.Customer
	if err := c.BodyParser(&updateCustomerId); err != nil || updateCustomerId.CustomerName == "" {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Valid Customer name is required",
		})
	}

	Customer.CustomerName = updateCustomerId.CustomerName
	db.DB.Save(&Customer)

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Update Name successfully",
		"data":    Customer,
	})
}

func DeleteCustomer(c *fiber.Ctx) error {
	CustomerId := c.Params("CustomerID")
	var Customer models.Customer

	if err := db.DB.First(&Customer, CustomerId).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "ID is Not Found",
		})
	}

	db.DB.Delete(&Customer)

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "ID is Deleted successfully",
	})
}
