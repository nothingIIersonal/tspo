package services

import (
	"tspo_final/internal/dtos"
	"tspo_final/internal/models"
	"tspo_final/internal/repository"
)

func CreateFeature(feature *models.Feature, repository repository.FeatureRepository) dtos.Response {
	operationResult := repository.Save(feature)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var data = operationResult.Result.(*models.Feature)

	return dtos.Response{Success: true, Data: data}
}

func FindAllFeatures(repository repository.FeatureRepository) dtos.Response {
	operationResult := repository.FindAll()

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var datas = operationResult.Result.(*models.Features)

	return dtos.Response{Success: true, Data: datas}
}

func FindOneFeatureById(id uint, repository repository.FeatureRepository) dtos.Response {
	operationResult := repository.FindOneById(id)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var data = operationResult.Result.(*models.Feature)

	return dtos.Response{Success: true, Data: data}
}

func UpdateFeatureById(id uint, feature *models.Feature, repository repository.FeatureRepository) dtos.Response {
	existingFeatureResponse := FindOneFeatureById(id, repository)

	if !existingFeatureResponse.Success {
		return existingFeatureResponse
	}

	existingFeature := existingFeatureResponse.Data.(*models.Feature)

	existingFeature.Feature = feature.Feature
	existingFeature.IsDeleted = feature.IsDeleted

	operationResult := repository.Save(existingFeature)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	return dtos.Response{Success: true, Data: operationResult.Result}
}

func DeleteOneFeatureById(id uint, repository repository.FeatureRepository) dtos.Response {
	operationResult := repository.DeleteOneById(id)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	return dtos.Response{Success: true}
}

func DeleteFeatureByIds(multiId *dtos.MultiID, repository repository.FeatureRepository) dtos.Response {
	operationResult := repository.DeleteByIds(&multiId.Ids)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	return dtos.Response{Success: true}
}
