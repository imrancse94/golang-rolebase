package routes

import (
	"go-fiber/src/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	// Setup Public Routes
	SetupPublicRoutes(app)

	// Setup Auth Middleware
	app.Use(middleware.AuthMiddleware)
	// Setup Auth Routes
	SetupAuthRoutes(app)

	// Setup User Routes
	SetupUserRoutes(app)

	// Setup User Group Routes
	SetupUserGroupRoutes(app)
}
