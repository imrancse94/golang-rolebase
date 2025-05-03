package services

import (
	"go-fiber/src/config"
	"go-fiber/src/models"
	"net/http"
)

func CreateUserGroup(userGroup *models.UserGroup) ServiceResponse {
	// Check if user group name exists
	userGroupModel := &models.UserGroup{}
	exists, _ := userGroupModel.NameExists(userGroup.Name)

	if exists {
		return NewServiceResponse(config.AlreadyExistsCode, http.StatusConflict, "User group name already in use", nil)
	}

	// Save user group to the database
	userGroupNew := models.UserGroup{
		Name: userGroup.Name,
	}

	if err := userGroupNew.Create(); err != nil {
		return NewServiceResponse(config.FailedCode, http.StatusInternalServerError, "Failed to create user group", nil)
	}

	return NewServiceResponse(config.SuccessCode, http.StatusCreated, "User group created successfully", userGroupNew)
}

func GetUserGroupByID(id string) ServiceResponse {
	userGroupModel := &models.UserGroup{}
	err := userGroupModel.FindByID(id)

	if err != nil {
		return NewServiceResponse(config.NotFoundCode, http.StatusNotFound, "User group not found", nil)
	}

	return NewServiceResponse(config.SuccessCode, http.StatusOK, "User group retrieved successfully", userGroupModel)
}

func UpdateUserGroup(id string, userGroup *models.UserGroup) ServiceResponse {
	// Check if user group name exists
	userGroupModel := &models.UserGroup{}
	exists, _ := userGroupModel.NameExists(userGroup.Name)

	if exists {
		return NewServiceResponse(config.AlreadyExistsCode, http.StatusConflict, "User group name already in use", nil)
	}

	// Update user group in the database
	userGroupNew := models.UserGroup{
		Name: userGroup.Name,
	}

	if err := userGroupNew.Update(id); err != nil {
		return NewServiceResponse(config.FailedCode, http.StatusInternalServerError, "Failed to update user group", nil)
	}

	return NewServiceResponse(config.SuccessCode, http.StatusOK, "User group updated successfully", userGroupNew)
}

func DeleteUserGroup(id string) ServiceResponse {
	userGroupModel := &models.UserGroup{}
	err := userGroupModel.FindByID(id)

	if err != nil {
		return NewServiceResponse(config.NotFoundCode, http.StatusNotFound, "User group not found", nil)
	}

	if err := userGroupModel.Delete(id); err != nil {
		return NewServiceResponse(config.FailedCode, http.StatusInternalServerError, "Failed to delete user group", nil)
	}

	return NewServiceResponse(config.SuccessCode, http.StatusOK, "User group deleted successfully", nil)
}

func GetAllUserGroups() ServiceResponse {
	userGroupModel := &models.UserGroup{}
	userGroups, err := userGroupModel.GetAll()

	if err != nil {
		return NewServiceResponse(config.FailedCode, http.StatusInternalServerError, "Failed to retrieve user groups", nil)
	}

	return NewServiceResponse(config.SuccessCode, http.StatusOK, "User groups retrieved successfully", userGroups)
}

func ListUserGroups() ServiceResponse {
	userGroupModel := &models.UserGroup{}
	userGroups, err := userGroupModel.GetAll()

	if err != nil {
		return NewServiceResponse(config.FailedCode, http.StatusInternalServerError, "Failed to retrieve user groups", nil)
	}

	return NewServiceResponse(config.SuccessCode, http.StatusOK, "User groups retrieved successfully", userGroups)
}

func AddUserToGroup(userID, groupID uint) ServiceResponse {
	userGroupModel := &models.UsergroupUser{
		UserID:      userID,
		UserGroupID: groupID,
	}
	err := userGroupModel.Create()

	if err != nil {
		return NewServiceResponse(config.FailedCode, http.StatusInternalServerError, "Failed to add user to group", nil)
	}

	return NewServiceResponse(config.SuccessCode, http.StatusOK, "User added to group successfully", nil)
}

func RemoveUserFromGroup(userID uint) ServiceResponse {
	userGroupModel := &models.UsergroupUser{
		UserID: userID,
	}
	err := userGroupModel.Delete()

	if err != nil {
		return NewServiceResponse(config.FailedCode, http.StatusInternalServerError, "Failed to remove user from group", nil)
	}

	return NewServiceResponse(config.SuccessCode, http.StatusOK, "User removed from group successfully", nil)
}
