package models

import (
	"context"
	"go-fiber/src/config"
)

type UsergroupUser struct {
	UserID      uint `json:"user_id"`
	UserGroupID uint `json:"usergroup_id"`
}

func (uug *UsergroupUser) Create() error {
	// execption handling
	ctx := context.Background()

	// Insert user group into the database
	return config.DB.Create(ctx, "usergroup_users", uug)
}

func (uug *UsergroupUser) Delete() error {
	// execption handling
	ctx := context.Background()

	// Delete user group from the database
	return config.DB.Delete(ctx, "usergroup_users", uug)
}
