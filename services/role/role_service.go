package services

import (
	"pr8_1/dtos"
	"pr8_1/models"
	"pr8_1/repositories"
)

func CreateRole(role *models.Role, repository repositories.RoleRepository) dtos.Response {
	operationResult := repository.Save(role)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var data = operationResult.Result.(*models.Role)

	return dtos.Response{Success: true, Data: data}
}

func FindAllRoles(repository repositories.RoleRepository) dtos.Response {
	operationResult := repository.FindAll()

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var datas = operationResult.Result.(*models.Roles)

	return dtos.Response{Success: true, Data: datas}
}

func FindOneRoleById(id uint, repository repositories.RoleRepository) dtos.Response {
	operationResult := repository.FindOneById(id)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var data = operationResult.Result.(*models.Role)

	return dtos.Response{Success: true, Data: data}
}

func UpdateRoleById(id uint, role *models.Role, repository repositories.RoleRepository) dtos.Response {
	existingRoleResponse := FindOneRoleById(id, repository)

	if !existingRoleResponse.Success {
		return existingRoleResponse
	}

	existingRole := existingRoleResponse.Data.(*models.Role)

	existingRole.Role = role.Role

	operationResult := repository.Save(existingRole)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	return dtos.Response{Success: true, Data: operationResult.Result}
}

func DeleteOneRoleById(id uint, repository repositories.RoleRepository) dtos.Response {
	operationResult := repository.DeleteOneById(id)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	return dtos.Response{Success: true}
}

func DeleteRoleByIds(multiId *dtos.MultiID, repository repositories.RoleRepository) dtos.Response {
	operationResult := repository.DeleteByIds(&multiId.Ids)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	return dtos.Response{Success: true}
}
