package internal

import (
	"os"

	"github.com/flambra/bank/internal/account"
	"github.com/flambra/bank/internal/institution"
	"github.com/flambra/bank/internal/transaction"
	"github.com/gofiber/fiber/v2"
)

func InitializeRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"project":     os.Getenv("PROJECT"),
			"environment": os.Getenv("ENV"),
			"version":     os.Getenv("BUILD_VERSION"),
		})
	})

	app.Post("/institution", institution.Create)
	app.Get("/institution/:id", institution.Read)
	app.Put("/institution/:id", institution.Update)
	app.Delete("/institution/:id", institution.Delete)
	app.Get("/institutions", institution.List)

	app.Post("/account", account.Create)
	app.Get("/account/:id", account.Read)
	app.Put("/account/:id", account.Update)
	app.Delete("/account/:id", account.Delete)

	app.Post("/transaction", transaction.Create)
	app.Get("/transactions", transaction.Page)

}
