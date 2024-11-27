package configs

import (
	contextlib "context"
	"net/http"
	"pr8_1/dtos"
	"pr8_1/models"
	"pr8_1/repositories"
	services "pr8_1/services/order"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func SetupOrdersRoutes(orderRepository *repositories.OrderRepository, route *gin.Engine) {
	// create route /create endpoint
	route.POST("/orders/create", func(context *gin.Context) {
		// initialization order model
		var order models.Order

		// validate json
		err := context.ShouldBindJSON(&order)

		// validation errors
		if err != nil {
			context.JSON(http.StatusBadRequest, "Can't bind JSON to order")
			return
		}

		// default http status code = 200
		code := http.StatusOK

		// save order & get it's response
		response := services.CreateOrder(&order, *orderRepository)

		// save order failed
		if !response.Success {
			// change http status code to 400
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	route.GET("/orders/", func(context *gin.Context) {
		code := http.StatusOK

		response := services.FindAllOrders(*orderRepository)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	route.GET("/orders/timeout/", func(context *gin.Context) {
		ctx, cancel := contextlib.WithTimeout(context.Request.Context(), 2*time.Second)
		defer cancel()

		code := http.StatusOK

		response := services.FindAllOrdersWithCtx(*orderRepository, &ctx)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	route.GET("/orders/paging/", func(context *gin.Context) {
		code := http.StatusOK

		page := 1
		limit := 10
		sort := "user_id asc"

		var searchs []dtos.Search

		query := context.Request.URL.Query()

		parseQuery(&query, &limit, &page, &sort, &searchs)

		offset := (page - 1) * limit

		response := services.FindAllOrdersPaging(*orderRepository, page, limit, offset, sort, searchs)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	route.GET("/orders/show/:id", func(context *gin.Context) {
		id_, _ := strconv.ParseUint(context.Param("id"), 10, 32)
		id := uint(id_)

		code := http.StatusOK

		response := services.FindOneOrderById(id, *orderRepository)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	route.PUT("/orders/update/:id", func(context *gin.Context) {
		id_, _ := strconv.ParseUint(context.Param("id"), 10, 32)
		id := uint(id_)

		var order models.Order

		err := context.ShouldBindJSON(&order)

		// validation errors
		if err != nil {
			context.JSON(http.StatusBadRequest, "Can't bind JSON to order")
			return
		}

		code := http.StatusOK

		response := services.UpdateOrderById(id, &order, *orderRepository)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	route.DELETE("/orders/delete/:id", func(context *gin.Context) {
		id_, _ := strconv.ParseUint(context.Param("id"), 10, 32)
		id := uint(id_)

		code := http.StatusOK

		response := services.DeleteOneOrderById(id, *orderRepository)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	route.POST("/orders/delete", func(context *gin.Context) {
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

		response := services.DeleteOrderByIds(&multiID, *orderRepository)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})
}
