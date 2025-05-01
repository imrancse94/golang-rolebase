package seeders

import (
	"go-fiber/src/config"
	"log"

	"gorm.io/gorm"
)

func TruncateAllTables(db *gorm.DB) error {
	var tableNames []string

	// Get all table names
	err := db.Raw("SHOW TABLES").Scan(&tableNames).Error
	if err != nil {
		return err
	}

	// Disable foreign key checks
	if err := db.Exec("SET FOREIGN_KEY_CHECKS = 0").Error; err != nil {
		return err
	}

	for _, table := range tableNames {
		if err := db.Exec("TRUNCATE TABLE " + table).Error; err != nil {
			return err
		}
	}

	// Re-enable foreign key checks
	if err := db.Exec("SET FOREIGN_KEY_CHECKS = 1").Error; err != nil {
		return err
	}

	return nil
}

func SeedAll(db *gorm.DB) {
	// Truncate tables before seeding
	TruncateAllTables(db)
	// Seed initial data
	if err := SeedUsers(db); err != nil {
		log.Println("Seeding Users failed:", err)
	}

	if err := PermissionSeeder(db); err != nil {
		log.Println("Seeding Permissions failed:", err)
	}

	if err := RoleSeeder(db); err != nil {
		log.Println("Seeding RoleSeeder failed:", err)
	}

	if err := UsergroupSeeder(db); err != nil {
		log.Println("Seeding UsergroupSeeder failed:", err)
	}

	if err := UsergroupRoleSeeder(db); err != nil {
		log.Println("Seeding UsergroupRoleSeeder failed:", err)
	}

	if err := RolePermissionSeeder(db); err != nil {
		log.Println("Seeding RolePermissionSeeder failed:", err)
	}

	if err := UsergroupUserSeeder(db); err != nil {
		log.Println("Seeding UsergroupUserSeeder failed:", err)
	}

}

func Run() {
	// Initialize the database connection
	dbIface, err := config.ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Run the seeders
	SeedAll(dbIface.SQLDB)
}
