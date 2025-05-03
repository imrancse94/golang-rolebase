package routes

import (
	"go-fiber/src/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupUserGroupRoutes(app *fiber.App) {
	userGroup := app.Group("/usergroups")
	userGroup.Get("/", controllers.ListUserGroups)
	userGroup.Post("/", controllers.CreateUserGroup)
	userGroup.Get("/:id", controllers.GetUserGroupByID)
	userGroup.Put("/:id", controllers.UpdateUserGroup)
	userGroup.Delete("/:id", controllers.DeleteUserGroup)
}
