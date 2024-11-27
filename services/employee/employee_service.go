package services

import (
	"pr8_1/dtos"
	"pr8_1/models"
	"pr8_1/repositories"
)

func CreateEmployee(employee *models.Employee, repository repositories.EmployeeRepository) dtos.Response {
	operationResult := repository.Save(employee)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var data = operationResult.Result.(*models.Employee)

	return dtos.Response{Success: true, Data: data}
}

func FindAllEmployees(repository repositories.EmployeeRepository) dtos.Response {
	operationResult := repository.FindAll()

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var datas = operationResult.Result.(*models.Employees)

	return dtos.Response{Success: true, Data: datas}
}

func FindOneEmployeeById(id uint, repository repositories.EmployeeRepository) dtos.Response {
	operationResult := repository.FindOneById(id)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var data = operationResult.Result.(*models.Employee)

	return dtos.Response{Success: true, Data: data}
}

func UpdateEmployeeById(id uint, employee *models.Employee, repository repositories.EmployeeRepository) dtos.Response {
	existingEmployeeResponse := FindOneEmployeeById(id, repository)

	if !existingEmployeeResponse.Success {
		return existingEmployeeResponse
	}

	existingEmployee := existingEmployeeResponse.Data.(*models.Employee)

	existingEmployee.Salary = employee.Salary
	existingEmployee.Position = employee.Position
	existingEmployee.KPI = employee.KPI

	operationResult := repository.Save(existingEmployee)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	return dtos.Response{Success: true, Data: operationResult.Result}
}

func DeleteOneEmployeeById(id uint, repository repositories.EmployeeRepository) dtos.Response {
	operationResult := repository.DeleteOneById(id)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	return dtos.Response{Success: true}
}

func DeleteEmployeeByIds(multiId *dtos.MultiID, repository repositories.EmployeeRepository) dtos.Response {
	operationResult := repository.DeleteByIds(&multiId.Ids)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	return dtos.Response{Success: true}
}
