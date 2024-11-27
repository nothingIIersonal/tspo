package routes

import (
	"net/http"
	"strconv"
	"tspo_final/internal/dtos"
	"tspo_final/internal/models"
	"tspo_final/internal/repository"
	services "tspo_final/internal/services/feature"

	"github.com/gin-gonic/gin"
)

var featureRepository *repository.FeatureRepository

// @Summary Create feature
// @Description Creates a features
// @Tags features
// @Accept json
// @Produce json
// @Param feature body models.Feature true "Feature details"
// @Success 201 {object} dtos.Response
// @Failure 400 {object} dtos.Response
// @Router /features/create [post]
func createFeature(context *gin.Context) {
	var feature models.Feature

	err := context.ShouldBindJSON(&feature)

	if err != nil {
		context.JSON(http.StatusBadRequest, "Can't bind JSON to feature")
		return
	}

	code := http.StatusOK

	response := services.CreateFeature(&feature, *featureRepository)

	if !response.Success {
		code = http.StatusBadRequest
	}

	context.JSON(code, response)
}

// @Summary Get all features
// @Description Returns list of features
// @Tags features
// @Produce json
// @Success 200 {object} dtos.Response
// @Failure 400 {object} dtos.Response
// @Router /features/ [get]
func getFeatures(context *gin.Context) {
	code := http.StatusOK

	response := services.FindAllFeatures(*featureRepository)

	if !response.Success {
		code = http.StatusBadRequest
	}

	context.JSON(code, response)
}

// @Summary Get one feature
// @Description Returns one feature
// @Tags features
// @Produce json
// @Param id path string true "Feature ID"
// @Success 200 {object} dtos.Response
// @Failure 400 {object} dtos.Response
// @Router /features/show/{id} [get]
func getFeature(context *gin.Context) {
	id_, _ := strconv.ParseUint(context.Param("id"), 10, 32)
	id := uint(id_)

	code := http.StatusOK

	response := services.FindOneFeatureById(id, *featureRepository)

	if !response.Success {
		code = http.StatusBadRequest
	}

	context.JSON(code, response)
}

// @Summary Update a feature
// @Description Updates a features
// @Tags features
// @Accept json
// @Produce json
// @Param id path string true "Feature ID"
// @Param feature body models.Feature true "Feature details"
// @Success 201 {object} dtos.Response
// @Failure 400 {object} dtos.Response
// @Router /features/update/{id} [put]
func updateFeature(context *gin.Context) {
	id_, _ := strconv.ParseUint(context.Param("id"), 10, 32)
	id := uint(id_)

	var feature models.Feature

	err := context.ShouldBindJSON(&feature)

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
}

// @Summary Delete a feature
// @Description Deletes a features
// @Tags features
// @Accept json
// @Produce json
// @Param id path string true "Feature ID"
// @Success 201 {object} dtos.Response
// @Failure 400 {object} dtos.Response
// @Router /features/delete/{id} [delete]
func deleteFeature(context *gin.Context) {
	id_, _ := strconv.ParseUint(context.Param("id"), 10, 32)
	id := uint(id_)

	code := http.StatusOK

	response := services.DeleteOneFeatureById(id, *featureRepository)

	if !response.Success {
		code = http.StatusBadRequest
	}

	context.JSON(code, response)
}

// @Summary Delete features
// @Description Deletes features
// @Tags features
// @Accept json
// @Produce json
// @Param ids body dtos.MultiID true "Feature IDs"
// @Success 201 {object} dtos.Response
// @Failure 400 {object} dtos.Response
// @Router /features/delete/ [post]
func deleteFeatures(context *gin.Context) {
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

	response := services.DeleteFeatureByIds(&multiID, *featureRepository)

	if !response.Success {
		code = http.StatusBadRequest
	}

	context.JSON(code, response)
}

func SetupFeaturesRoutes(featureRepository_ *repository.FeatureRepository, route *gin.Engine) {
	featureRepository = featureRepository_
	route.POST("/features/create", createFeature)
	route.GET("/features/", getFeatures)
	route.GET("/features/show/:id", getFeature)
	route.PUT("/features/update/:id", updateFeature)
	route.DELETE("/features/delete/:id", deleteFeature)
	route.POST("/features/delete", deleteFeatures)
}
