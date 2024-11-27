package services

import (
	"context"
	"pr8_1/dtos"
	"pr8_1/models"
	"pr8_1/repositories"
)

func CreateUser(user *models.User, repository repositories.UserRepository) dtos.Response {
	operationResult := repository.Save(user)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var data = operationResult.Result.(*models.User)

	return dtos.Response{Success: true, Data: data}
}

func FindAllUsers(repository repositories.UserRepository) dtos.Response {
	operationResult := repository.FindAll()

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var datas = operationResult.Result.(*models.Users)

	return dtos.Response{Success: true, Data: datas}
}

func FindAllUsersWithCtx(repository repositories.UserRepository, ctx *context.Context) dtos.Response {
	operationResult := repository.FindAllWithCtx(ctx)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var datas = operationResult.Result.(*models.Users)

	return dtos.Response{Success: true, Data: datas}
}

func FindAllUsersPaging(repository repositories.UserRepository, page int, limit int, offset int, sort string, searchs []dtos.Search) dtos.ResponsePaging {
	operationResult, total := repository.FindAllPaging(limit, offset, sort, searchs)

	if operationResult.Error != nil {
		return dtos.ResponsePaging{Success: false, Message: operationResult.Error.Error()}
	}

	var datas = operationResult.Result.(*models.Users)

	return dtos.ResponsePaging{Success: true, Data: datas, Total: &total, Page: page, Limit: &limit}
}

func FindOneUserById(id uint, repository repositories.UserRepository) dtos.Response {
	operationResult := repository.FindOneById(id)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var data = operationResult.Result.(*models.User)

	return dtos.Response{Success: true, Data: data}
}

func UpdateUserById(id uint, user *models.User, repository repositories.UserRepository) dtos.Response {
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

func DeleteOneUserById(id uint, repository repositories.UserRepository) dtos.Response {
	operationResult := repository.DeleteOneById(id)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	return dtos.Response{Success: true}
}

func DeleteUserByIds(multiId *dtos.MultiID, repository repositories.UserRepository) dtos.Response {
	operationResult := repository.DeleteByIds(&multiId.Ids)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	return dtos.Response{Success: true}
}
