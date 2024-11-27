package configs

import (
	contextlib "context"
	"net/http"
	"pr8_1/dtos"
	"pr8_1/models"
	"pr8_1/repositories"
	services "pr8_1/services/good"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func SetupGoodsRoutes(goodRepository *repositories.GoodRepository, route *gin.Engine) {
	// create route /create endpoint
	route.POST("/goods/create", func(context *gin.Context) {
		// initialization good model
		var good models.Good

		// validate json
		err := context.ShouldBindJSON(&good)

		// validation errors
		if err != nil {
			context.JSON(http.StatusBadRequest, "Can't bind JSON to good")
			return
		}

		// default http status code = 200
		code := http.StatusOK

		// save good & get it's response
		response := services.CreateGood(&good, *goodRepository)

		// save good failed
		if !response.Success {
			// change http status code to 400
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	route.GET("/goods/", func(context *gin.Context) {
		code := http.StatusOK

		response := services.FindAllGoods(*goodRepository)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	route.GET("/goods/timeout/", func(context *gin.Context) {
		ctx, cancel := contextlib.WithTimeout(context.Request.Context(), 2*time.Second)
		defer cancel()

		code := http.StatusOK

		response := services.FindAllGoodsWithCtx(*goodRepository, &ctx)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	route.GET("/goods/paging/", func(context *gin.Context) {
		code := http.StatusOK

		page := 1
		limit := 10
		sort := "user_id asc"

		var searchs []dtos.Search

		query := context.Request.URL.Query()

		parseQuery(&query, &limit, &page, &sort, &searchs)

		offset := (page - 1) * limit

		response := services.FindAllGoodsPaging(*goodRepository, page, limit, offset, sort, searchs)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	route.GET("/goods/show/:id", func(context *gin.Context) {
		id_, _ := strconv.ParseUint(context.Param("id"), 10, 32)
		id := uint(id_)

		code := http.StatusOK

		response := services.FindOneGoodById(id, *goodRepository)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	route.PUT("/goods/update/:id", func(context *gin.Context) {
		id_, _ := strconv.ParseUint(context.Param("id"), 10, 32)
		id := uint(id_)

		var good models.Good

		err := context.ShouldBindJSON(&good)

		// validation errors
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
	})

	route.DELETE("/goods/delete/:id", func(context *gin.Context) {
		id_, _ := strconv.ParseUint(context.Param("id"), 10, 32)
		id := uint(id_)

		code := http.StatusOK

		response := services.DeleteOneGoodById(id, *goodRepository)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	route.POST("/goods/delete", func(context *gin.Context) {
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

		response := services.DeleteGoodByIds(&multiID, *goodRepository)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})
}
