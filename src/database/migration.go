package database

import (
	"fmt"
	"go-fiber/src/config"
	"go-fiber/src/database/migrations"
	"os"
)

func AutoMigrate(db *config.Database) error {

	dbType := os.Getenv("DB_CONNECTION")

	switch dbType {

	case "mysql", "postgres":

		db.SQLDB.AutoMigrate(
			migrations.Entities...,
		)

	case "mongodb":

	default:
		return fmt.Errorf("unsupported database type: %s", dbType)
	}

	return nil
}
