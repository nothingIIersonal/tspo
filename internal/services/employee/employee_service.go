package services

import (
	"tspo_final/internal/dtos"
	"tspo_final/internal/models"
	"tspo_final/internal/repository"
)

func CreateEmployee(employee *models.Employee, repository repository.EmployeeRepository) dtos.Response {
	operationResult := repository.Save(employee)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var data = operationResult.Result.(*models.Employee)

	return dtos.Response{Success: true, Data: data}
}

func FindAllEmployees(repository repository.EmployeeRepository) dtos.Response {
	operationResult := repository.FindAll()

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var datas = operationResult.Result.(*models.Employees)

	return dtos.Response{Success: true, Data: datas}
}

func FindOneEmployeeById(id uint, repository repository.EmployeeRepository) dtos.Response {
	operationResult := repository.FindOneById(id)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var data = operationResult.Result.(*models.Employee)

	return dtos.Response{Success: true, Data: data}
}

func UpdateEmployeeById(id uint, employee *models.Employee, repository repository.EmployeeRepository) dtos.Response {
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

func DeleteOneEmployeeById(id uint, repository repository.EmployeeRepository) dtos.Response {
	operationResult := repository.DeleteOneById(id)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	return dtos.Response{Success: true}
}

func DeleteEmployeeByIds(multiId *dtos.MultiID, repository repository.EmployeeRepository) dtos.Response {
	operationResult := repository.DeleteByIds(&multiId.Ids)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	return dtos.Response{Success: true}
}
