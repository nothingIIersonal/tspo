package services

import (
	"tspo_final/internal/dtos"
	"tspo_final/internal/models"
	"tspo_final/internal/repository"
)

func CreateGood(good *models.Good, repository repository.GoodRepository) dtos.Response {
	operationResult := repository.Save(good)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var data = operationResult.Result.(*models.Good)

	return dtos.Response{Success: true, Data: data}
}

func FindAllGoods(repository repository.GoodRepository) dtos.Response {
	operationResult := repository.FindAll()

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var datas = operationResult.Result.(*models.Goods)

	return dtos.Response{Success: true, Data: datas}
}

func FindOneGoodById(id uint, repository repository.GoodRepository) dtos.Response {
	operationResult := repository.FindOneById(id)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var data = operationResult.Result.(*models.Good)

	return dtos.Response{Success: true, Data: data}
}

func UpdateGoodById(id uint, good *models.Good, repository repository.GoodRepository) dtos.Response {
	existingGoodResponse := FindOneGoodById(id, repository)

	if !existingGoodResponse.Success {
		return existingGoodResponse
	}

	existingGood := existingGoodResponse.Data.(*models.Good)

	existingGood.Name = good.Name
	existingGood.Description = good.Description
	existingGood.Price = good.Price
	existingGood.Count = good.Count
	existingGood.IsDeleted = good.IsDeleted

	operationResult := repository.Save(existingGood)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	return dtos.Response{Success: true, Data: operationResult.Result}
}

func DeleteOneGoodById(id uint, repository repository.GoodRepository) dtos.Response {
	operationResult := repository.DeleteOneById(id)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	return dtos.Response{Success: true}
}

func DeleteGoodByIds(multiId *dtos.MultiID, repository repository.GoodRepository) dtos.Response {
	operationResult := repository.DeleteByIds(&multiId.Ids)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	return dtos.Response{Success: true}
}
