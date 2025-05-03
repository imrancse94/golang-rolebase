package services

import (
	"fmt"
	"go-fiber/src/config"
	"go-fiber/src/lib"
	"go-fiber/src/models"
	"net/http"
)

func Register(user *models.User) ServiceResponse {
	// Check if email exists
	userModel := &models.User{}
	exists, _ := userModel.EmailExists(user.Email)

	fmt.Println("Email exists:", exists)
	if exists {
		return NewServiceResponse(config.EmailExistsCode, http.StatusConflict, "Email already in use", nil)
	}

	// Hash the password
	hashedPassword, err := lib.HashPassword(user.Password)
	if err != nil {
		return NewServiceResponse(config.FailedCode, http.StatusInternalServerError, "Failed to hash password", nil)
	}
	user.Password = hashedPassword

	// Save user to the database
	userNew := models.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}

	if err := userNew.Create(); err != nil {
		return NewServiceResponse(config.FailedCode, http.StatusInternalServerError, "Failed to create user", nil)
	}

	return NewServiceResponse(config.SuccessCode, http.StatusCreated, "User registered successfully", userNew)
}

func Login(user *models.User) ServiceResponse {

	// Check if user exists
	userModel := &models.User{}
	err := userModel.FindByEmail(user.Email)

	if err != nil {
		return NewServiceResponse(config.NotFoundCode, http.StatusNotFound, "User not found", nil)
	}

	// Check password
	if isValidPassword := lib.CheckPasswordHash(user.Password, userModel.Password); !isValidPassword {
		return NewServiceResponse(config.UnauthorizedCode, http.StatusUnauthorized, "Invalid credentials", nil)
	}

	payload := map[string]interface{}{
		"email": userModel.Email,
		"id":    userModel.ID,
		"name":  userModel.Name,
	}

	// Generate tokens
	token, err := lib.GenerateToken(payload)
	if err != nil {
		return NewServiceResponse(config.FailedCode, http.StatusInternalServerError, "Failed to generate token", nil)
	}

	refreshToken, err := lib.GenerateRefreshToken(payload)
	if err != nil {
		return NewServiceResponse(config.FailedCode, http.StatusInternalServerError, "Failed to generate refresh token", nil)
	}

	permissions := GetPermissionsByUserID(userModel.ID)
	// Success response
	data := map[string]interface{}{
		"access_token":  token,
		"refresh_token": refreshToken,
		"user":          userModel,
		"permissions":   permissions,
	}

	return NewServiceResponse(config.SuccessCode, http.StatusOK, "Login successful", data)
}

func GetPermissionsByUserID(userID uint) []models.Permission {
	// Fetch permissions from the database
	permissionModel := &models.Permission{}
	permissions, err := permissionModel.GetPermissionsByUserID(userID)

	if err != nil {
		return []models.Permission{}
	}

	return permissions
}
