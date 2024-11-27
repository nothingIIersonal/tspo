package routes

import (
	"net/http"
	"strconv"
	"tspo_final/internal/dtos"
	"tspo_final/internal/models"
	"tspo_final/internal/repository"
	services "tspo_final/internal/services/employee"

	"github.com/gin-gonic/gin"
)

var employeeRepository *repository.EmployeeRepository

// @Summary Create employee
// @Description Creates a employees
// @Tags employees
// @Accept json
// @Produce json
// @Param employee body models.Employee true "Employee details"
// @Success 201 {object} dtos.Response
// @Failure 400 {object} dtos.Response
// @Router /employees/create [post]
func createEmployee(context *gin.Context) {
	var employee models.Employee

	err := context.ShouldBindJSON(&employee)

	if err != nil {
		context.JSON(http.StatusBadRequest, "Can't bind JSON to employee")
		return
	}

	code := http.StatusOK

	response := services.CreateEmployee(&employee, *employeeRepository)

	if !response.Success {
		code = http.StatusBadRequest
	}

	context.JSON(code, response)
}

// @Summary Get all employees
// @Description Returns list of employees
// @Tags employees
// @Produce json
// @Success 200 {object} dtos.Response
// @Failure 400 {object} dtos.Response
// @Router /employees/ [get]
func getEmployees(context *gin.Context) {
	code := http.StatusOK

	response := services.FindAllEmployees(*employeeRepository)

	if !response.Success {
		code = http.StatusBadRequest
	}

	context.JSON(code, response)
}

// @Summary Get one employee
// @Description Returns one employee
// @Tags employees
// @Produce json
// @Param id path string true "Employee ID"
// @Success 200 {object} dtos.Response
// @Failure 400 {object} dtos.Response
// @Router /employees/show/{id} [get]
func getEmployee(context *gin.Context) {
	id_, _ := strconv.ParseUint(context.Param("id"), 10, 32)
	id := uint(id_)

	code := http.StatusOK

	response := services.FindOneEmployeeById(id, *employeeRepository)

	if !response.Success {
		code = http.StatusBadRequest
	}

	context.JSON(code, response)
}

// @Summary Update a employee
// @Description Updates a employees
// @Tags employees
// @Accept json
// @Produce json
// @Param id path string true "Employee ID"
// @Param employee body models.Employee true "Employee details"
// @Success 201 {object} dtos.Response
// @Failure 400 {object} dtos.Response
// @Router /employees/update/{id} [put]
func updateEmployee(context *gin.Context) {
	id_, _ := strconv.ParseUint(context.Param("id"), 10, 32)
	id := uint(id_)

	var employee models.Employee

	err := context.ShouldBindJSON(&employee)

	if err != nil {
		context.JSON(http.StatusBadRequest, "Can't bind JSON to employee")
		return
	}

	code := http.StatusOK

	response := services.UpdateEmployeeById(id, &employee, *employeeRepository)

	if !response.Success {
		code = http.StatusBadRequest
	}

	context.JSON(code, response)
}

// @Summary Delete a employee
// @Description Deletes a employees
// @Tags employees
// @Accept json
// @Produce json
// @Param id path string true "Employee ID"
// @Success 201 {object} dtos.Response
// @Failure 400 {object} dtos.Response
// @Router /employees/delete/{id} [delete]
func deleteEmployee(context *gin.Context) {
	id_, _ := strconv.ParseUint(context.Param("id"), 10, 32)
	id := uint(id_)

	code := http.StatusOK

	response := services.DeleteOneEmployeeById(id, *employeeRepository)

	if !response.Success {
		code = http.StatusBadRequest
	}

	context.JSON(code, response)
}

// @Summary Delete employees
// @Description Deletes employees
// @Tags employees
// @Accept json
// @Produce json
// @Param ids body dtos.MultiID true "Employee IDs"
// @Success 201 {object} dtos.Response
// @Failure 400 {object} dtos.Response
// @Router /employees/delete/ [post]
func deleteEmployees(context *gin.Context) {
	var multiID dtos.MultiID

	err := context.ShouldBindJSON(&multiID)

	if err != nil {
		context.JSON(http.StatusBadRequest, "Can't bind JSON to multiID")
		return
	}

	if len(multiID.Ids) == 0 {
		response := dtos.Response{Success: false, Message: "IDs cannot be empty."}

		context.JSON(http.StatusBadRequest, response)

		return
	}

	code := http.StatusOK

	response := services.DeleteEmployeeByIds(&multiID, *employeeRepository)

	if !response.Success {
		code = http.StatusBadRequest
	}

	context.JSON(code, response)
}

func SetupEmployeesRoutes(employeeRepository_ *repository.EmployeeRepository, route *gin.Engine) {
	employeeRepository = employeeRepository_
	route.POST("/employees/create", createEmployee)
	route.GET("/employees/", getEmployees)
	route.GET("/employees/show/:id", getEmployee)
	route.PUT("/employees/update/:id", updateEmployee)
	route.DELETE("/employees/delete/:id", deleteEmployee)
	route.POST("/employees/delete", deleteEmployees)
}
