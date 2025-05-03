package models

import (
	"context"
	"go-fiber/src/config"
	"go-fiber/src/lib"
)

type Permission struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Key  string `json:"key"`
}

func (p *Permission) GetPermissionsByUserID(userID uint) ([]Permission, error) {
	ctx := context.Background()
	permissions := []Permission{}

	// Query database for permissions by user ID
	builder := lib.NewQueryBuilder().
		Select("p.*").
		From("users u").
		Join("usergroup_users ugu", "ugu.user_id = u.id").
		Join("usergroup_roles ugr", "ugr.usergroup_id = ugu.usergroup_id").
		Join("role_permissions rp", "rp.role_id = ugr.role_id").
		Join("permissions p", "p.id = rp.permission_id").
		Where("u.id = ?", userID)

	// execute the query
	err := config.DB.QueryWithBuilder(ctx, builder, &permissions)

	if err != nil {
		return nil, err
	}

	return permissions, nil
}
