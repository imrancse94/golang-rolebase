package requests

type UserGroupCreateRequest struct {
	Name string `json:"name" validate:"required,min=3"`
}
type UserGroupUpdateRequest struct {
	Name string `json:"name" validate:"required,min=3"`
}
type UserGroupDeleteRequest struct {
	ID string `json:"id" validate:"required"`
}
type UserGroupGetRequest struct {
	ID string `json:"id" validate:"required"`
}
type UserGroupGetAllRequest struct {
	ID string `json:"id" validate:"required"`
}

type UserGroupAddUserRequest struct {
	UserID      uint `json:"user_id" validate:"required"`
	UserGroupID uint `json:"usergroup_id" validate:"required"`
}

type UserGroupRemoveUserRequest struct {
	UserID uint `json:"user_id" validate:"required"`
}
