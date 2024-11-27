package services

import (
	"pr8_1/dtos"
	"pr8_1/models"
	"pr8_1/repositories"
)

func CreateOrder(order *models.Order, repository repositories.OrderRepository) dtos.Response {
	operationResult := repository.Save(order)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var data = operationResult.Result.(*models.Order)

	return dtos.Response{Success: true, Data: data}
}

func FindAllOrders(repository repositories.OrderRepository) dtos.Response {
	operationResult := repository.FindAll()

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var datas = operationResult.Result.(*models.Orders)

	return dtos.Response{Success: true, Data: datas}
}

func FindOneOrderById(id uint, repository repositories.OrderRepository) dtos.Response {
	operationResult := repository.FindOneById(id)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var data = operationResult.Result.(*models.Order)

	return dtos.Response{Success: true, Data: data}
}

func UpdateOrderById(id uint, order *models.Order, repository repositories.OrderRepository) dtos.Response {
	existingOrderResponse := FindOneOrderById(id, repository)

	if !existingOrderResponse.Success {
		return existingOrderResponse
	}

	existingOrder := existingOrderResponse.Data.(*models.Order)

	existingOrder.DeliveryType = order.DeliveryType
	existingOrder.DeliveryTime = order.DeliveryTime
	existingOrder.OrderTime = order.OrderTime
	existingOrder.TotalPrice = order.TotalPrice
	existingOrder.Canceled = order.Canceled

	operationResult := repository.Save(existingOrder)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	return dtos.Response{Success: true, Data: operationResult.Result}
}

func DeleteOneOrderById(id uint, repository repositories.OrderRepository) dtos.Response {
	operationResult := repository.DeleteOneById(id)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	return dtos.Response{Success: true}
}

func DeleteOrderByIds(multiId *dtos.MultiID, repository repositories.OrderRepository) dtos.Response {
	operationResult := repository.DeleteByIds(&multiId.Ids)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	return dtos.Response{Success: true}
}
