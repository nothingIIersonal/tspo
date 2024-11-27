package routes

import (
	"net/http"
	"strconv"
	"tspo_final/internal/dtos"
	"tspo_final/internal/models"
	"tspo_final/internal/repository"
	services "tspo_final/internal/services/role"

	"github.com/gin-gonic/gin"
)

func SetupRolesRoutes(roleRepository *repository.RoleRepository, route *gin.Engine) {
	route.POST("/roles/create", func(context *gin.Context) {
		var role models.Role

		err := context.ShouldBindJSON(&role)

		if err != nil {
			context.JSON(http.StatusBadRequest, "Can't bind JSON to role")
			return
		}

		code := http.StatusOK

		response := services.CreateRole(&role, *roleRepository)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	route.GET("/roles/", func(context *gin.Context) {
		code := http.StatusOK

		response := services.FindAllRoles(*roleRepository)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	route.GET("/roles/show/:id", func(context *gin.Context) {
		id_, _ := strconv.ParseUint(context.Param("id"), 10, 32)
		id := uint(id_)

		code := http.StatusOK

		response := services.FindOneRoleById(id, *roleRepository)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	route.PUT("/roles/update/:id", func(context *gin.Context) {
		id_, _ := strconv.ParseUint(context.Param("id"), 10, 32)
		id := uint(id_)

		var role models.Role

		err := context.ShouldBindJSON(&role)

		if err != nil {
			context.JSON(http.StatusBadRequest, "Can't bind JSON to role")
			return
		}

		code := http.StatusOK

		response := services.UpdateRoleById(id, &role, *roleRepository)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	route.DELETE("/roles/delete/:id", func(context *gin.Context) {
		id_, _ := strconv.ParseUint(context.Param("id"), 10, 32)
		id := uint(id_)

		code := http.StatusOK

		response := services.DeleteOneRoleById(id, *roleRepository)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	route.POST("/roles/delete", func(context *gin.Context) {
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

		response := services.DeleteRoleByIds(&multiID, *roleRepository)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})
}
