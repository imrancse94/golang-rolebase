package config

import (
	"context"
	"errors"
	"fmt"
	"go-fiber/src/database/migrations"
	"go-fiber/src/lib"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB lib.DBInterface

type Database struct {
	SQLDB   *gorm.DB
	NoSQLDB *mongo.Database
}

// ConnectDB initializes the database connection based on DB_CONNECTION type
func ConnectDB() (*Database, error) {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using system environment variables")
	}

	dbType := os.Getenv("DB_CONNECTION")

	dbIface := &Database{}

	switch dbType {
	case "mysql":
		dsn := fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_DATABASE"),
		)

		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			return nil, fmt.Errorf("failed to connect to MySQL: %v", err)
		}
		DB = &lib.SQLDB{Db: db}
		dbIface.SQLDB = db
		db.AutoMigrate(migrations.Entities...)

	case "postgres":
		dsn := fmt.Sprintf(
			"user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_DATABASE"),
		)
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			return nil, fmt.Errorf("failed to connect to PostgreSQL: %v", err)
		}
		DB = &lib.SQLDB{Db: db}
		dbIface.SQLDB = db
		db.AutoMigrate(migrations.Entities...)
	case "mongodb":
		client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(os.Getenv("MONGODB_URI")))
		if err != nil {
			return nil, fmt.Errorf("failed to connect to MongoDB: %v", err)
		}
		db := client.Database(os.Getenv("MONGODB_NAME"))
		DB = &lib.MongoDB{Db: db}
		dbIface.NoSQLDB = db

	default:
		return nil, errors.New("unsupported database type: " + dbType)
	}

	return dbIface, nil
}
