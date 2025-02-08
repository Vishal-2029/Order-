package routes

import (
	"Exercise/OrderAPI/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	//authentication routes
	app.Post("/Customer/:CustomerId/login", controllers.Login)
	app.Post("/Customer/:CustomerId/logout", controllers.Logout)
	app.Get("/Customer/:Customerd/passcode", controllers.Passcode)

	//customer routes...
	app.Post("/customer", controllers.CreateCustomers)
	app.Get("/customer", controllers.CustomerList)
	app.Put("/customer/:CustomerId", controllers.UpdateCustomer)
	app.Delete("/customer/:CustomerId", controllers.DeleteCustomer)

	//Items routes...
	app.Post("/items", controllers.CreateItems)
	app.Get("/items", controllers.ItemList)
	app.Put("/items/:ItemId", controllers.UpdateItems)
	app.Delete("/items/:ItemId", controllers.DeleteItem)

	//Order routes...
	app.Post("/Order", controllers.PlaceOrder)
	app.Put("/Order/:CustomerId/approve", controllers.ApproveOrder)
	app.Put("/Order/:CustomerId/cancel", controllers.CancelOrder)
	app.Get("/Order/:CustomerId/bill", controllers.GetBill)

}
