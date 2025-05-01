package routes

import (
	"go-fiber/src/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupAuthRoutes(app *fiber.App) {
	authGroup := app.Group("/auth")
	authGroup.Post("/login", controllers.Login)
}
