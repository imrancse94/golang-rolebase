package utils

import "github.com/gofiber/fiber/v2"

// fiber response common methods
func SuccessResponse(c *fiber.Ctx, httpCode int, statusCode int, message string, data any) error {
	return c.Status(httpCode).JSON(fiber.Map{
		"status_code": statusCode,
		"message":     message,
		"data":        data,
	})
}

func ErrorResponse(c *fiber.Ctx, httpCode int, statusCode int, message string) error {
	return c.Status(httpCode).JSON(fiber.Map{
		"status_code": statusCode,
		"message":     message,
	})
}

func ValidationErrorResponse(c *fiber.Ctx, httpCode int, statusCode int, message string, errors map[string]string) error {
	return c.Status(httpCode).JSON(fiber.Map{
		"status_code": statusCode,
		"message":     message,
		"errors":      errors,
	})
}
