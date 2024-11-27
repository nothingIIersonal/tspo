package routes

import (
	"net/http"
	"strconv"
	"tspo_final/internal/dtos"
	"tspo_final/internal/models"
	"tspo_final/internal/repository"
	services "tspo_final/internal/services/good"

	"github.com/gin-gonic/gin"
)

var goodRepository *repository.GoodRepository

// @Summary Create good
// @Description Creates a goods
// @Tags goods
// @Accept json
// @Produce json
// @Param good body models.Good true "Good details"
// @Success 201 {object} dtos.Response
// @Failure 400 {object} dtos.Response
// @Router /goods/create [post]
func createGood(context *gin.Context) {
	var good models.Good

	err := context.ShouldBindJSON(&good)

	if err != nil {
		context.JSON(http.StatusBadRequest, "Can't bind JSON to good")
		return
	}

	code := http.StatusOK

	response := services.CreateGood(&good, *goodRepository)

	if !response.Success {
		code = http.StatusBadRequest
	}

	context.JSON(code, response)
}

// @Summary Get all goods
// @Description Returns list of goods
// @Tags goods
// @Produce json
// @Success 200 {object} dtos.Response
// @Failure 400 {object} dtos.Response
// @Router /goods/ [get]
func getGoods(context *gin.Context) {
	code := http.StatusOK

	response := services.FindAllGoods(*goodRepository)

	if !response.Success {
		code = http.StatusBadRequest
	}

	context.JSON(code, response)
}

// @Summary Get one good
// @Description Returns one good
// @Tags goods
// @Produce json
// @Param id path string true "Good ID"
// @Success 200 {object} dtos.Response
// @Failure 400 {object} dtos.Response
// @Router /goods/show/{id} [get]
func getGood(context *gin.Context) {
	id_, _ := strconv.ParseUint(context.Param("id"), 10, 32)
	id := uint(id_)

	code := http.StatusOK

	response := services.FindOneGoodById(id, *goodRepository)

	if !response.Success {
		code = http.StatusBadRequest
	}

	context.JSON(code, response)
}

// @Summary Update a good
// @Description Updates a goods
// @Tags goods
// @Accept json
// @Produce json
// @Param id path string true "Good ID"
// @Param good body models.Good true "Good details"
// @Success 201 {object} dtos.Response
// @Failure 400 {object} dtos.Response
// @Router /goods/update/{id} [put]
func updateGood(context *gin.Context) {
	id_, _ := strconv.ParseUint(context.Param("id"), 10, 32)
	id := uint(id_)

	var good models.Good

	err := context.ShouldBindJSON(&good)

	if err != nil {
		context.JSON(http.StatusBadRequest, "Can't bind JSON to good")
		return
	}

	code := http.StatusOK

	response := services.UpdateGoodById(id, &good, *goodRepository)

	if !response.Success {
		code = http.StatusBadRequest
	}

	context.JSON(code, response)
}

// @Summary Delete a good
// @Description Deletes a goods
// @Tags goods
// @Accept json
// @Produce json
// @Param id path string true "Good ID"
// @Success 201 {object} dtos.Response
// @Failure 400 {object} dtos.Response
// @Router /goods/delete/{id} [delete]
func deleteGood(context *gin.Context) {
	id_, _ := strconv.ParseUint(context.Param("id"), 10, 32)
	id := uint(id_)

	code := http.StatusOK

	response := services.DeleteOneGoodById(id, *goodRepository)

	if !response.Success {
		code = http.StatusBadRequest
	}

	context.JSON(code, response)
}

// @Summary Delete goods
// @Description Deletes goods
// @Tags goods
// @Accept json
// @Produce json
// @Param ids body dtos.MultiID true "Good IDs"
// @Success 201 {object} dtos.Response
// @Failure 400 {object} dtos.Response
// @Router /goods/delete/ [post]
func deleteGoods(context *gin.Context) {
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

	response := services.DeleteGoodByIds(&multiID, *goodRepository)

	if !response.Success {
		code = http.StatusBadRequest
	}

	context.JSON(code, response)
}

func SetupGoodsRoutes(goodRepository_ *repository.GoodRepository, route *gin.Engine) {
	goodRepository = goodRepository_
	route.POST("/goods/create", createGood)
	route.GET("/goods/", getGoods)
	route.GET("/goods/show/:id", getGood)
	route.PUT("/goods/update/:id", updateGood)
	route.DELETE("/goods/delete/:id", deleteGood)
	route.POST("/goods/delete", deleteGoods)
}
