package services

import (
	"tspo_final/internal/dtos"
	"tspo_final/internal/models"
	"tspo_final/internal/repository"
)

func CreateUser(user *models.User, repository repository.UserRepository) dtos.Response {
	operationResult := repository.Save(user)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var data = operationResult.Result.(*models.User)

	return dtos.Response{Success: true, Data: data}
}

func FindAllUsers(repository repository.UserRepository) dtos.Response {
	operationResult := repository.FindAll()

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var datas = operationResult.Result.(*models.Users)

	return dtos.Response{Success: true, Data: datas}
}

func FindOneUserById(id uint, repository repository.UserRepository) dtos.Response {
	operationResult := repository.FindOneById(id)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var data = operationResult.Result.(*models.User)

	return dtos.Response{Success: true, Data: data}
}

func UpdateUserById(id uint, user *models.User, repository repository.UserRepository) dtos.Response {
	existingUserResponse := FindOneUserById(id, repository)

	if !existingUserResponse.Success {
		return existingUserResponse
	}

	existingUser := existingUserResponse.Data.(*models.User)

	existingUser.Name = user.Name
	existingUser.Phone = user.Phone
	existingUser.Address = user.Address
	existingUser.Email = user.Email
	existingUser.PasswordHash = user.PasswordHash
	existingUser.IsDeleted = user.IsDeleted

	operationResult := repository.Save(existingUser)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	return dtos.Response{Success: true, Data: operationResult.Result}
}

func DeleteOneUserById(id uint, repository repository.UserRepository) dtos.Response {
	operationResult := repository.DeleteOneById(id)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	return dtos.Response{Success: true}
}

func DeleteUserByIds(multiId *dtos.MultiID, repository repository.UserRepository) dtos.Response {
	operationResult := repository.DeleteByIds(&multiId.Ids)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	return dtos.Response{Success: true}
}
