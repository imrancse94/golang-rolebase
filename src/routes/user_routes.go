package routes

import (
	"go-fiber/src/controllers"
	"go-fiber/src/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	userGroup := app.Group("/users")
	userGroup.Post("/", controllers.CreateUser)
	userGroup.Get("/:id", middleware.AuthMiddleware, controllers.GetUserByID) // Protected Route
}
