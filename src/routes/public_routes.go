package routes

import (
	"go-fiber/src/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupPublicRoutes(app *fiber.App) {
	// Public Routes
	public := app.Group("/")
	public.Post("/login", controllers.Login)
}
