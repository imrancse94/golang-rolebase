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

func CreateUser(c *fiber.Ctx) error {
	var req requests.UserCreateRequest

	// Parse and validate request body
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	// Validate request
	if errors := lib.ValidateStruct(req); errors != nil {
		return utils.ValidationErrorResponse(c, fiber.StatusUnprocessableEntity, config.ValidationCode, "Validation error", errors)
	}

	response := services.Register(&models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	})

	if response.StatusCode == config.SuccessCode {

		return utils.SuccessResponse(c, response.HTTPCode, response.StatusCode, response.Message, response.Data)
	}

	return c.Status(response.HTTPCode).JSON(fiber.Map{
		"error": response.Message,
	})

}

func GetUserByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	if err := user.FindByID(id); err != nil {
		return c.Status(404).SendString("User not found")
	}
	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	var req requests.UserLoginRequest

	// Parse and validate request body
	if err := c.BodyParser(&req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, config.FailedCode, "Invalid request")
	}

	// Validate request
	if errors := lib.ValidateStruct(req); errors != nil {
		return utils.ValidationErrorResponse(c, fiber.StatusUnprocessableEntity, config.ValidationCode, "Validation error", errors)
	}

	response := services.Login(&models.User{
		Email:    req.Email,
		Password: req.Password,
	})

	if response.StatusCode == config.SuccessCode {
		return utils.SuccessResponse(c, response.HTTPCode, response.StatusCode, response.Message, response.Data)
	}

	return utils.ErrorResponse(c, response.HTTPCode, response.StatusCode, response.Message)

}
