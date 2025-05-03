package models

import (
	"context"
	"go-fiber/src/config"
)

type UserGroup struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func (ug *UserGroup) Create() error {
	ctx := context.Background()

	// Insert user group into the database
	return config.DB.Create(ctx, "usergroups", ug)
}

func (ug *UserGroup) Update(id string) error {
	ctx := context.Background()

	// Update user group in the database
	return config.DB.Update(ctx, "usergroups", id, ug)
}
func (ug *UserGroup) FindByID(id string) error {
	ctx := context.Background()
	existingUserGroup := new(UserGroup)

	// Query database for existing user group
	err := config.DB.Find(ctx, "usergroups", map[string]interface{}{"id": id}, existingUserGroup)
	if err != nil {
		return err
	}

	// If user group is found, return it
	*ug = *existingUserGroup
	return nil
}
func (ug *UserGroup) NameExists(name string) (bool, error) {
	ctx := context.Background()
	existingUserGroup := new(UserGroup)

	// Query database for existing user group name
	err := config.DB.Find(ctx, "usergroups", map[string]interface{}{"name": name}, existingUserGroup)
	if err != nil {
		return false, err
	}

	// If user group name is found, return true
	return existingUserGroup.Name != "", nil
}
func (ug *UserGroup) Delete(id string) error {
	ctx := context.Background()

	// Delete user group from the database
	return config.DB.Delete(ctx, "usergroups", id)
}
func (ug *UserGroup) GetAll() ([]UserGroup, error) {
	ctx := context.Background()
	var userGroups []UserGroup

	// Query database for all user groups
	err := config.DB.FindAll(ctx, "usergroups", nil, &userGroups)
	if err != nil {
		return nil, err
	}

	// Return the list of user groups
	return userGroups, nil
}
