package configs

import (
	contextlib "context"
	"net/http"
	"pr8_1/dtos"
	"pr8_1/models"
	"pr8_1/repositories"
	services "pr8_1/services/vendor"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func SetupVendorsRoutes(vendorRepository *repositories.VendorRepository, route *gin.Engine) {
	// create route /create endpoint
	route.POST("/vendors/create", func(context *gin.Context) {
		// initialization vendor model
		var vendor models.Vendor

		// validate json
		err := context.ShouldBindJSON(&vendor)

		// validation errors
		if err != nil {
			context.JSON(http.StatusBadRequest, "Can't bind JSON to vendor")
			return
		}

		// default http status code = 200
		code := http.StatusOK

		// save vendor & get it's response
		response := services.CreateVendor(&vendor, *vendorRepository)

		// save vendor failed
		if !response.Success {
			// change http status code to 400
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	route.GET("/vendors/", func(context *gin.Context) {
		code := http.StatusOK

		response := services.FindAllVendors(*vendorRepository)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	route.GET("/vendors/timeout/", func(context *gin.Context) {
		ctx, cancel := contextlib.WithTimeout(context.Request.Context(), 2*time.Second)
		defer cancel()

		code := http.StatusOK

		response := services.FindAllVendorsWithCtx(*vendorRepository, &ctx)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	route.GET("/vendors/paging/", func(context *gin.Context) {
		code := http.StatusOK

		page := 1
		limit := 10
		sort := "user_id asc"

		var searchs []dtos.Search

		query := context.Request.URL.Query()

		parseQuery(&query, &limit, &page, &sort, &searchs)

		offset := (page - 1) * limit

		response := services.FindAllVendorsPaging(*vendorRepository, page, limit, offset, sort, searchs)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	route.GET("/vendors/show/:id", func(context *gin.Context) {
		id_, _ := strconv.ParseUint(context.Param("id"), 10, 32)
		id := uint(id_)

		code := http.StatusOK

		response := services.FindOneVendorById(id, *vendorRepository)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	route.PUT("/vendors/update/:id", func(context *gin.Context) {
		id_, _ := strconv.ParseUint(context.Param("id"), 10, 32)
		id := uint(id_)

		var vendor models.Vendor

		err := context.ShouldBindJSON(&vendor)

		// validation errors
		if err != nil {
			context.JSON(http.StatusBadRequest, "Can't bind JSON to vendor")
			return
		}

		code := http.StatusOK

		response := services.UpdateVendorById(id, &vendor, *vendorRepository)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	route.DELETE("/vendors/delete/:id", func(context *gin.Context) {
		id_, _ := strconv.ParseUint(context.Param("id"), 10, 32)
		id := uint(id_)

		code := http.StatusOK

		response := services.DeleteOneVendorById(id, *vendorRepository)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	route.POST("/vendors/delete", func(context *gin.Context) {
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

		response := services.DeleteVendorByIds(&multiID, *vendorRepository)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})
}
