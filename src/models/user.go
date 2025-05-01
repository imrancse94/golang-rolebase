package models

import (
	"context"
	"go-fiber/src/config"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

// HashPassword hashes the password before storing it
func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// Check if the email already exists in the database
func (u *User) EmailExists(email string) (bool, error) {
	ctx := context.Background()
	existingUser := new(User)

	// Query database for existing email
	err := config.DB.Find(ctx, "users", map[string]interface{}{"email": email}, existingUser)
	if err != nil {
		return false, err
	}

	// If email is found, return true
	return existingUser.Email != "", nil
}

// Create inserts a new user with a hashed password
func (u *User) Create() error {
	ctx := context.Background()

	// Hash the password before saving
	if err := u.HashPassword(); err != nil {
		return err
	}

	return config.DB.Create(ctx, "users", u)
}

// Update modifies a user and hashes the password if changed
func (u *User) Update() error {
	ctx := context.Background()

	// Hash password only if it's changed
	if u.Password != "" {
		if err := u.HashPassword(); err != nil {
			return err
		}
	}

	return config.DB.Update(ctx, "users", map[string]interface{}{"id": u.ID}, u)
}

func (u *User) FindByID(id string) error {
	ctx := context.Background()
	return config.DB.Find(ctx, "users", map[string]interface{}{"id": id}, u)
}

func (u *User) Delete() error {
	ctx := context.Background()
	return config.DB.Delete(ctx, "users", map[string]interface{}{"id": u.ID})
}

// FindByEmail retrieves a user by their email from the database
func (u *User) FindByEmail(email string) error {
	ctx := context.Background()

	// Query the database for the user with the given email
	return config.DB.Find(ctx, "users", map[string]interface{}{"email": email}, u)
}
