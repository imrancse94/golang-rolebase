package seeders

import (
	"strings"
	"unicode"

	"go-fiber/src/database/migrations"

	"gorm.io/gorm"
)

// Capitalize first letter
func capitalize(str string) string {
	if str == "" {
		return ""
	}
	runes := []rune(str)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

func PermissionSeeder(db *gorm.DB) error {
	managementModules := []string{
		"Role Management",
		"User Management",
		"Usergroup Management",
		"Usergroup Role Management",
	}

	actions := []string{"list", "create", "update", "delete", "view"}

	for _, moduleName := range managementModules {
		// Generate key prefix (e.g., "Usergroup-role")
		prefix := strings.ToLower(strings.ReplaceAll(strings.TrimSuffix(moduleName, " Management"), " ", "-"))

		// Create parent permission
		parent := migrations.Permission{
			Name:     moduleName,
			ParentID: 0,
		}

		if err := db.Create(&parent).Error; err != nil {
			return err
		}

		// Create child permissions
		for _, action := range actions {
			capitalName := capitalize(action)
			key := prefix + "-" + action

			child := migrations.Permission{
				Name:     capitalName,
				ParentID: parent.ID,
				Key:      key,
			}

			if err := db.Create(&child).Error; err != nil {
				return err
			}
		}
	}

	return nil
}

func RoleSeeder(db *gorm.DB) error {
	// Create roles
	roles := []migrations.Role{
		{Name: "Admin"},
		{Name: "User"},
	}

	for _, role := range roles {
		if err := db.Create(&role).Error; err != nil {
			return err
		}
	}

	return nil
}

func UsergroupSeeder(db *gorm.DB) error {
	// Create user groups
	Usergroups := []migrations.Usergroup{
		{Name: "Admin"},
		{Name: "User"},
	}

	for _, group := range Usergroups {
		if err := db.Create(&group).Error; err != nil {
			return err
		}
	}

	return nil
}

func UsergroupRoleSeeder(db *gorm.DB) error {
	// Create user group roles
	UsergroupRoles := []migrations.UsergroupRole{
		{UsergroupID: 1, RoleID: 1},
		// {UsergroupID: 2, RoleID: 2},
	}

	for _, groupRole := range UsergroupRoles {
		if err := db.Create(&groupRole).Error; err != nil {
			return err
		}
	}

	return nil
}
func UsergroupUserSeeder(db *gorm.DB) error {
	// Create user group users
	UsergroupUsers := []migrations.UsergroupUser{
		{UserID: 1, UsergroupID: 1},
		// {UserID: 2, UsergroupID: 2},
	}

	for _, groupUser := range UsergroupUsers {
		if err := db.Create(&groupUser).Error; err != nil {
			return err
		}
	}

	return nil
}

func RolePermissionSeeder(db *gorm.DB) error {
	var permissions []migrations.Permission

	// Get all child permissions (ParentID != 0)
	if err := db.Where("parent_id != 0").Find(&permissions).Error; err != nil {
		return err
	}

	// Insert into role_permissions table
	for _, permission := range permissions {
		isIndex := false
		if permission.Key != "" && strings.Contains(permission.Key, "list") {
			isIndex = true
		}
		rp := migrations.RolePermission{
			RoleID:       1,
			PermissionID: permission.ID,
			IsIndex:      isIndex,
		}

		if err := db.Create(&rp).Error; err != nil {
			return err
		}
	}

	return nil
}
