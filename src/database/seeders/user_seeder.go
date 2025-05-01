package seeders

import (
	"go-fiber/src/database/migrations"

	"gorm.io/gorm"
)

func SeedUsers(db *gorm.DB) error {
	user := migrations.User{
		Name:     "Imran Hossain",
		Email:    "imrancse94@gmail.com",
		Password: "$2a$12$3Mqsmflq5/9niV9oYFeph.LEsz/5AitKJG661YpjBBx6W2Lss6isu", // bcrypted password for "Nop@ss1234"
	}
	if err := db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}
