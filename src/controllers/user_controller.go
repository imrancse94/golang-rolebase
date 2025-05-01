package controllers

import (
	"go-fiber/src/lib"
	"go-fiber/src/models"
	"go-fiber/src/requests"
	"os"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(c *fiber.Ctx) error {
	var req requests.UserCreateRequest

	// Parse and validate request body
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	// Validate request
	if errors := lib.ValidateStruct(req); errors != nil {
		return c.Status(422).JSON(errors) // Return validation errors
	}

	// Check if email already exists
	userModel := &models.User{}
	exists, _ := userModel.EmailExists(req.Email)

	if exists {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": "Email already in use",
		})
	}

	// Create user
	user := models.User{Name: req.Name, Email: req.Email, Password: req.Password}
	if err := user.Create(); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create user"})
	}

	return c.Status(201).JSON(user)
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
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	// Validate request
	if errors := lib.ValidateStruct(req); errors != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(errors) // Return validation errors
	}

	// Find user by email
	user := new(models.User)
	err := user.FindByEmail(req.Email)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	// Compare the provided password with the stored hashed password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	// Prepare the payload with dynamic data
	payload := map[string]interface{}{
		"email": user.Email,
		"id":    user.ID,
		"name":  user.Name,
	}

	// Generate JWT access token
	accessToken, err := lib.GenerateToken(payload)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error generating access token"})
	}

	// Generate JWT refresh token
	refreshToken, err := lib.GenerateRefreshToken(payload)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error generating refresh token"})
	}

	return c.JSON(fiber.Map{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
		"expires_in":    os.Getenv("JWT_ACCESS_TOKEN_EXPIRES_IN"), // in hour
		"user":          user,
	})
}
