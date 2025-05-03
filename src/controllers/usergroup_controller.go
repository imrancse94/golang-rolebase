package controllers

import (
	"go-fiber/src/config"
	"go-fiber/src/lib"
	"go-fiber/src/models"
	"go-fiber/src/requests"
	"go-fiber/src/services"
	"go-fiber/src/utils"

	"github.com/gofiber/fiber/v2"
)

func CreateUserGroup(c *fiber.Ctx) error {
	var req requests.UserGroupCreateRequest

	// Parse and validate request body
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	// Validate request
	if errors := lib.ValidateStruct(req); errors != nil {
		return utils.ValidationErrorResponse(c, fiber.StatusUnprocessableEntity, config.ValidationCode, "Validation error", errors)
	}

	response := services.CreateUserGroup(&models.UserGroup{
		Name: req.Name,
	})

	if response.StatusCode == config.SuccessCode {

		return utils.SuccessResponse(c, response.HTTPCode, response.StatusCode, response.Message, response.Data)
	}

	return c.Status(response.HTTPCode).JSON(fiber.Map{
		"error": response.Message,
	})
}

func GetUserGroupByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var userGroup models.UserGroup
	if err := userGroup.FindByID(id); err != nil {
		return c.Status(404).SendString("User group not found")
	}
	return c.JSON(userGroup)
}

func UpdateUserGroup(c *fiber.Ctx) error {
	id := c.Params("id")
	var req requests.UserGroupUpdateRequest

	// Parse and validate request body
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	// Validate request
	if errors := lib.ValidateStruct(req); errors != nil {
		return utils.ValidationErrorResponse(c, fiber.StatusUnprocessableEntity, config.ValidationCode, "Validation error", errors)
	}

	response := services.UpdateUserGroup(id, &models.UserGroup{
		Name: req.Name,
	})

	if response.StatusCode == config.SuccessCode {

		return utils.SuccessResponse(c, response.HTTPCode, response.StatusCode, response.Message, response.Data)
	}

	return c.Status(response.HTTPCode).JSON(fiber.Map{
		"error": response.Message,
	})
}

func DeleteUserGroup(c *fiber.Ctx) error {
	id := c.Params("id")
	response := services.DeleteUserGroup(id)

	if response.StatusCode == config.SuccessCode {

		return utils.SuccessResponse(c, response.HTTPCode, response.StatusCode, response.Message, response.Data)
	}

	return c.Status(response.HTTPCode).JSON(fiber.Map{
		"error": response.Message,
	})
}
func ListUserGroups(c *fiber.Ctx) error {
	response := services.ListUserGroups()

	if response.StatusCode == config.SuccessCode {

		return utils.SuccessResponse(c, response.HTTPCode, response.StatusCode, response.Message, response.Data)
	}

	return c.Status(response.HTTPCode).JSON(fiber.Map{
		"error": response.Message,
	})
}

func AddUserToGroup(c *fiber.Ctx) error {

	var req requests.UserGroupAddUserRequest

	// Parse and validate request body
	if err := c.BodyParser(&req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, config.FailedCode, "Invalid request")
	}

	// Validate request
	if errors := lib.ValidateStruct(req); errors != nil {
		return utils.ValidationErrorResponse(c, fiber.StatusUnprocessableEntity, config.ValidationCode, "Validation error", errors)
	}

	response := services.AddUserToGroup(req.UserID, req.UserGroupID)

	if response.StatusCode == config.SuccessCode {

		return utils.SuccessResponse(c, response.HTTPCode, response.StatusCode, response.Message, response.Data)
	}

	return utils.ErrorResponse(c, response.HTTPCode, response.StatusCode, response.Message)
}

func RemoveUserFromGroup(c *fiber.Ctx) error {

	var req requests.UserGroupRemoveUserRequest

	// Parse and validate request body
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	// Validate request
	if errors := lib.ValidateStruct(req); errors != nil {
		return utils.ValidationErrorResponse(c, fiber.StatusUnprocessableEntity, config.ValidationCode, "Validation error", errors)
	}

	response := services.RemoveUserFromGroup(req.UserID)

	if response.StatusCode == config.SuccessCode {

		return utils.SuccessResponse(c, response.HTTPCode, response.StatusCode, response.Message, response.Data)
	}

	return utils.ErrorResponse(c, response.HTTPCode, response.StatusCode, response.Message)
}
