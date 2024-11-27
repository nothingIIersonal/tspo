package configs

import (
	"net/http"
	"pr8_1/dtos"
	"pr8_1/models"
	"pr8_1/repositories"
	services "pr8_1/services/employee"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SetupEmployeesRoutes(employeeRepository *repositories.EmployeeRepository, route *gin.Engine) {
	// create route /create endpoint
	route.POST("/employees/create", func(context *gin.Context) {
		// initialization employee model
		var employee models.Employee

		// validate json
		err := context.ShouldBindJSON(&employee)

		// validation errors
		if err != nil {
			context.JSON(http.StatusBadRequest, "Can't bind JSON to employee")
			return
		}

		// default http status code = 200
		code := http.StatusOK

		// save employee & get it's response
		response := services.CreateEmployee(&employee, *employeeRepository)

		// save employee failed
		if !response.Success {
			// change http status code to 400
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	route.GET("/employees/", func(context *gin.Context) {
		code := http.StatusOK

		response := services.FindAllEmployees(*employeeRepository)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	route.GET("/employees/show/:id", func(context *gin.Context) {
		id_, _ := strconv.ParseUint(context.Param("id"), 10, 32)
		id := uint(id_)

		code := http.StatusOK

		response := services.FindOneEmployeeById(id, *employeeRepository)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	route.PUT("/employees/update/:id", func(context *gin.Context) {
		id_, _ := strconv.ParseUint(context.Param("id"), 10, 32)
		id := uint(id_)

		var employee models.Employee

		err := context.ShouldBindJSON(&employee)

		// validation errors
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
	})

	route.DELETE("/employees/delete/:id", func(context *gin.Context) {
		id_, _ := strconv.ParseUint(context.Param("id"), 10, 32)
		id := uint(id_)

		code := http.StatusOK

		response := services.DeleteOneEmployeeById(id, *employeeRepository)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	route.POST("/employees/delete", func(context *gin.Context) {
		var multiID dtos.MultiID

		err := context.ShouldBindJSON(&multiID)

		// validation errors
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
	})
}
