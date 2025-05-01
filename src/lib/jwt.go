package lib

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Secret key for JWT
var jwtSecret = []byte(os.Getenv("JWT_SECRET_KEY"))

// GenerateToken generates a JWT access token for the user with dynamic payload
func GenerateToken(payload map[string]interface{}) (string, error) {

	expiryValue, _ := strconv.Atoi(os.Getenv("JWT_ACCESS_TOKEN_EXPIRES_IN"))

	claims := jwt.MapClaims{
		"exp": time.Now().Add(time.Second * time.Duration(expiryValue)).Unix(), // Expire in 1 day
	}

	// Add the dynamic payload to the claims
	for key, value := range payload {
		claims[key] = value
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// GenerateRefreshToken generates a JWT refresh token for the user with dynamic payload
func GenerateRefreshToken(payload map[string]interface{}) (string, error) {

	expiryValue, _ := strconv.Atoi(os.Getenv("JWT_REFRESH_TOKEN_EXPIRES_IN"))

	claims := jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * time.Duration(expiryValue)).Unix(), // Expire in 30 days
	}

	// Add the dynamic payload to the claims
	for key, value := range payload {
		claims[key] = value
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// ValidateToken verifies the JWT token and returns claims
func ValidateToken(tokenString string) (jwt.MapClaims, error) {
	// Extract token from Bearer prefix
	if !strings.HasPrefix(tokenString, "Bearer ") {
		return nil, errors.New("Authorization header must start with Bearer")
	}

	// Remove the "Bearer " prefix from the token string
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	// Handle error if token is invalid or cannot be parsed
	if err != nil {
		return nil, err
	}

	// Extract claims from the token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("Invalid token")
	}

	// Optionally log token claims for debugging
	fmt.Println("Token claims:", claims)

	return claims, nil
}
