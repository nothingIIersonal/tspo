package configs

import (
	contextlib "context"
	"net/http"
	"pr8_1/dtos"
	"pr8_1/models"
	"pr8_1/repositories"
	services "pr8_1/services/role"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func SetupRolesRoutes(roleRepository *repositories.RoleRepository, route *gin.Engine) {
	// create route /create endpoint
	route.POST("/roles/create", func(context *gin.Context) {
		// initialization role model
		var role models.Role

		// validate json
		err := context.ShouldBindJSON(&role)

		// validation errors
		if err != nil {
			context.JSON(http.StatusBadRequest, "Can't bind JSON to role")
			return
		}

		// default http status code = 200
		code := http.StatusOK

		// save role & get it's response
		response := services.CreateRole(&role, *roleRepository)

		// save role failed
		if !response.Success {
			// change http status code to 400
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

	route.GET("/roles/timeout/", func(context *gin.Context) {
		ctx, cancel := contextlib.WithTimeout(context.Request.Context(), 2*time.Second)
		defer cancel()

		code := http.StatusOK

		response := services.FindAllRolesWithCtx(*roleRepository, &ctx)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	route.GET("/roles/paging/", func(context *gin.Context) {
		code := http.StatusOK

		page := 1
		limit := 10
		sort := "user_id asc"

		var searchs []dtos.Search

		query := context.Request.URL.Query()

		parseQuery(&query, &limit, &page, &sort, &searchs)

		offset := (page - 1) * limit

		response := services.FindAllRolesPaging(*roleRepository, page, limit, offset, sort, searchs)

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

		// validation errors
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

		response := services.DeleteRoleByIds(&multiID, *roleRepository)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})
}
