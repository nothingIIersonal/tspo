package configs

import (
	contextlib "context"
	"net/http"
	"pr8_1/dtos"
	"pr8_1/models"
	"pr8_1/repositories"
	services "pr8_1/services/feature"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func SetupFeaturesRoutes(featureRepository *repositories.FeatureRepository, route *gin.Engine) {
	// create route /create endpoint
	route.POST("/features/create", func(context *gin.Context) {
		// initialization feature model
		var feature models.Feature

		// validate json
		err := context.ShouldBindJSON(&feature)

		// validation errors
		if err != nil {
			context.JSON(http.StatusBadRequest, "Can't bind JSON to feature")
			return
		}

		// default http status code = 200
		code := http.StatusOK

		// save feature & get it's response
		response := services.CreateFeature(&feature, *featureRepository)

		// save feature failed
		if !response.Success {
			// change http status code to 400
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	route.GET("/features/", func(context *gin.Context) {
		code := http.StatusOK

		response := services.FindAllFeatures(*featureRepository)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	route.GET("/features/timeout/", func(context *gin.Context) {
		ctx, cancel := contextlib.WithTimeout(context.Request.Context(), 2*time.Second)
		defer cancel()

		code := http.StatusOK

		response := services.FindAllFeaturesWithCtx(*featureRepository, &ctx)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	route.GET("/features/paging/", func(context *gin.Context) {
		code := http.StatusOK

		page := 1
		limit := 10
		sort := "user_id asc"

		var searchs []dtos.Search

		query := context.Request.URL.Query()

		parseQuery(&query, &limit, &page, &sort, &searchs)

		offset := (page - 1) * limit

		response := services.FindAllFeaturesPaging(*featureRepository, page, limit, offset, sort, searchs)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	route.GET("/features/show/:id", func(context *gin.Context) {
		id_, _ := strconv.ParseUint(context.Param("id"), 10, 32)
		id := uint(id_)

		code := http.StatusOK

		response := services.FindOneFeatureById(id, *featureRepository)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	route.PUT("/features/update/:id", func(context *gin.Context) {
		id_, _ := strconv.ParseUint(context.Param("id"), 10, 32)
		id := uint(id_)

		var feature models.Feature

		err := context.ShouldBindJSON(&feature)

		// validation errors
		if err != nil {
			context.JSON(http.StatusBadRequest, "Can't bind JSON to feature")
			return
		}

		code := http.StatusOK

		response := services.UpdateFeatureById(id, &feature, *featureRepository)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	route.DELETE("/features/delete/:id", func(context *gin.Context) {
		id_, _ := strconv.ParseUint(context.Param("id"), 10, 32)
		id := uint(id_)

		code := http.StatusOK

		response := services.DeleteOneFeatureById(id, *featureRepository)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	route.POST("/features/delete", func(context *gin.Context) {
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

		response := services.DeleteFeatureByIds(&multiID, *featureRepository)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})
}
