package services

import (
	"context"
	"pr8_1/dtos"
	"pr8_1/models"
	"pr8_1/repositories"
)

func CreateVendor(vendor *models.Vendor, repository repositories.VendorRepository) dtos.Response {
	operationResult := repository.Save(vendor)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var data = operationResult.Result.(*models.Vendor)

	return dtos.Response{Success: true, Data: data}
}

func FindAllVendors(repository repositories.VendorRepository) dtos.Response {
	operationResult := repository.FindAll()

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var datas = operationResult.Result.(*models.Vendors)

	return dtos.Response{Success: true, Data: datas}
}

func FindAllVendorsWithCtx(repository repositories.VendorRepository, ctx *context.Context) dtos.Response {
	operationResult := repository.FindAllWithCtx(ctx)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var datas = operationResult.Result.(*models.Vendors)

	return dtos.Response{Success: true, Data: datas}
}

func FindAllVendorsPaging(repository repositories.VendorRepository, page int, limit int, offset int, sort string, searchs []dtos.Search) dtos.ResponsePaging {
	operationResult, total := repository.FindAllPaging(limit, offset, sort, searchs)

	if operationResult.Error != nil {
		return dtos.ResponsePaging{Success: false, Message: operationResult.Error.Error()}
	}

	var datas = operationResult.Result.(*models.Vendors)

	return dtos.ResponsePaging{Success: true, Data: datas, Total: &total, Page: page, Limit: &limit}
}

func FindOneVendorById(id uint, repository repositories.VendorRepository) dtos.Response {
	operationResult := repository.FindOneById(id)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var data = operationResult.Result.(*models.Vendor)

	return dtos.Response{Success: true, Data: data}
}

func UpdateVendorById(id uint, vendor *models.Vendor, repository repositories.VendorRepository) dtos.Response {
	existingVendorResponse := FindOneVendorById(id, repository)

	if !existingVendorResponse.Success {
		return existingVendorResponse
	}

	existingVendor := existingVendorResponse.Data.(*models.Vendor)

	existingVendor.Phone = vendor.Phone
	existingVendor.OrgName = vendor.OrgName
	existingVendor.INN = vendor.INN
	existingVendor.OGRN = vendor.OGRN
	existingVendor.Address = vendor.Address
	existingVendor.IsDeleted = vendor.IsDeleted

	operationResult := repository.Save(existingVendor)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	return dtos.Response{Success: true, Data: operationResult.Result}
}

func DeleteOneVendorById(id uint, repository repositories.VendorRepository) dtos.Response {
	operationResult := repository.DeleteOneById(id)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	return dtos.Response{Success: true}
}

func DeleteVendorByIds(multiId *dtos.MultiID, repository repositories.VendorRepository) dtos.Response {
	operationResult := repository.DeleteByIds(&multiId.Ids)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	return dtos.Response{Success: true}
}
