package controllers

import (
	db "Exercise/OrderAPI/config"
	"Exercise/OrderAPI/models"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func Login(c *fiber.Ctx) error {
	cashierId := c.Params("cashierId")
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"Message": "Invalid post request",
		})
	}
	if data["passcode"] == "" {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Passcode is required",
			"error":   map[string]interface{}{},
		})
	}
	var Customer models.Customer
	db.DB.Where("id = ?", cashierId).First(&Customer)

	if Customer.CustomerId == 0 {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Customer Not found",
			"error":   map[string]interface{}{},
		})
	}

	if Customer.Passcode != data["passcode"] {
		return c.Status(401).JSON(fiber.Map{
			"success": false,
			"message": "Passcode Not Match",
			"error":   map[string]interface{}{},
		})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Issuer":    strconv.Itoa(int(Customer.CustomerId)),
		"ExpiresAt": time.Now().Add(time.Hour * 24).Unix(), //1 day
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"success": false,
			"message": "Token Expired or invalid",
		})
	}

	CustomerData := make(map[string]interface{})
	CustomerData["token"] = tokenString

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Success",
		"data":    CustomerData,
	})

}
func Logout(c *fiber.Ctx) error {
	CustomerId := c.Params("CustomerId")
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	//check if passcode is empty
	if data["passcode"] == "" {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"Message": "Passcode is required",
		})
	}

	var Customer models.Customer
	db.DB.Where("Id = ?", CustomerId).First(&Customer)

	//check if cashier exist
	if Customer.CustomerId == 0 {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"Message": "Customer Not found",
		})
	}
	//check if passcode match
	if Customer.Passcode != data["passcode"] {
		return c.Status(401).JSON(fiber.Map{
			"success": false,
			"Message": "Passcode Not Match",
		})
	}
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)
	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"Message": "success",
	})
}
func Passcode(c *fiber.Ctx) error {
	CustomerId := c.Params("CustomerrId")
	var Customer models.Customer
	db.DB.Select("id,name,passcode").Where("id=?", CustomerId).First(&Customer)

	if Customer.CustomerName == "" || Customer.CustomerId == 0 {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Customer Not Found",
			"error":   map[string]interface{}{},
		})
	}

	CustomerData := make(map[string]interface{})
	CustomerData["passcode"] = Customer.Passcode

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Success",
		"data":    CustomerData,
	})
}
