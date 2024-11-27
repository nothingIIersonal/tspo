package routes

import (
	"net/http"
	"strconv"
	"tspo_final/internal/dtos"
	"tspo_final/internal/models"
	repository "tspo_final/internal/repository"
	services "tspo_final/internal/services/order"

	"github.com/gin-gonic/gin"
)

var orderRepository *repository.OrderRepository

// @Summary Create order
// @Description Creates a orders
// @Tags orders
// @Accept json
// @Produce json
// @Param order body models.Order true "Order details"
// @Success 201 {object} dtos.Response
// @Failure 400 {object} dtos.Response
// @Router /orders/create [post]
func createOrder(context *gin.Context) {
	var order models.Order

	err := context.ShouldBindJSON(&order)

	if err != nil {
		context.JSON(http.StatusBadRequest, "Can't bind JSON to order")
		return
	}

	code := http.StatusOK

	response := services.CreateOrder(&order, *orderRepository)

	if !response.Success {
		code = http.StatusBadRequest
	}

	context.JSON(code, response)
}

// @Summary Get all orders
// @Description Returns list of orders
// @Tags orders
// @Produce json
// @Success 200 {object} dtos.Response
// @Failure 400 {object} dtos.Response
// @Router /orders/ [get]
func getOrders(context *gin.Context) {
	code := http.StatusOK

	response := services.FindAllOrders(*orderRepository)

	if !response.Success {
		code = http.StatusBadRequest
	}

	context.JSON(code, response)
}

// @Summary Get one order
// @Description Returns one order
// @Tags orders
// @Produce json
// @Param id path string true "Order ID"
// @Success 200 {object} dtos.Response
// @Failure 400 {object} dtos.Response
// @Router /orders/show/{id} [get]
func getOrder(context *gin.Context) {
	id_, _ := strconv.ParseUint(context.Param("id"), 10, 32)
	id := uint(id_)

	code := http.StatusOK

	response := services.FindOneOrderById(id, *orderRepository)

	if !response.Success {
		code = http.StatusBadRequest
	}

	context.JSON(code, response)
}

// @Summary Update a order
// @Description Updates a orders
// @Tags orders
// @Accept json
// @Produce json
// @Param id path string true "Order ID"
// @Param order body models.Order true "Order details"
// @Success 201 {object} dtos.Response
// @Failure 400 {object} dtos.Response
// @Router /orders/update/{id} [put]
func updateOrder(context *gin.Context) {
	id_, _ := strconv.ParseUint(context.Param("id"), 10, 32)
	id := uint(id_)

	var order models.Order

	err := context.ShouldBindJSON(&order)

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
}

// @Summary Delete a order
// @Description Deletes a orders
// @Tags orders
// @Accept json
// @Produce json
// @Param id path string true "Order ID"
// @Success 201 {object} dtos.Response
// @Failure 400 {object} dtos.Response
// @Router /orders/delete/{id} [delete]
func deleteOrder(context *gin.Context) {
	id_, _ := strconv.ParseUint(context.Param("id"), 10, 32)
	id := uint(id_)

	code := http.StatusOK

	response := services.DeleteOneOrderById(id, *orderRepository)

	if !response.Success {
		code = http.StatusBadRequest
	}

	context.JSON(code, response)
}

// @Summary Delete orders
// @Description Deletes orders
// @Tags orders
// @Accept json
// @Produce json
// @Param ids body dtos.MultiID true "Order IDs"
// @Success 201 {object} dtos.Response
// @Failure 400 {object} dtos.Response
// @Router /orders/delete/ [post]
func deleteOrders(context *gin.Context) {
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

	response := services.DeleteOrderByIds(&multiID, *orderRepository)

	if !response.Success {
		code = http.StatusBadRequest
	}

	context.JSON(code, response)
}

func SetupOrdersRoutes(orderRepository_ *repository.OrderRepository, route *gin.Engine) {
	orderRepository = orderRepository_
	route.POST("/orders/create", createOrder)
	route.GET("/orders/", getOrders)
	route.GET("/orders/show/:id", getOrder)
	route.PUT("/orders/update/:id", updateOrder)
	route.DELETE("/orders/delete/:id", deleteOrder)
	route.POST("/orders/delete", deleteOrders)
}
