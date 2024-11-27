package routes

import (
	"net/http"
	"strconv"
	"tspo_final/internal/dtos"
	"tspo_final/internal/models"
	"tspo_final/internal/repository"
	services "tspo_final/internal/services/vendor"

	"github.com/gin-gonic/gin"
)

var vendorRepository *repository.VendorRepository

// @Summary Create vendor
// @Description Creates a vendors
// @Tags vendors
// @Accept json
// @Produce json
// @Param vendor body models.Vendor true "Vendor details"
// @Success 201 {object} dtos.Response
// @Failure 400 {object} dtos.Response
// @Router /vendors/create [post]
func createVendor(context *gin.Context) {
	var vendor models.Vendor

	err := context.ShouldBindJSON(&vendor)

	if err != nil {
		context.JSON(http.StatusBadRequest, "Can't bind JSON to vendor")
		return
	}

	code := http.StatusOK

	response := services.CreateVendor(&vendor, *vendorRepository)

	if !response.Success {
		code = http.StatusBadRequest
	}

	context.JSON(code, response)
}

// @Summary Get all vendors
// @Description Returns list of vendors
// @Tags vendors
// @Produce json
// @Success 200 {object} dtos.Response
// @Failure 400 {object} dtos.Response
// @Router /vendors/ [get]
func getVendors(context *gin.Context) {
	code := http.StatusOK

	response := services.FindAllVendors(*vendorRepository)

	if !response.Success {
		code = http.StatusBadRequest
	}

	context.JSON(code, response)
}

// @Summary Get one vendor
// @Description Returns one vendor
// @Tags vendors
// @Produce json
// @Param id path string true "Vendor ID"
// @Success 200 {object} dtos.Response
// @Failure 400 {object} dtos.Response
// @Router /vendors/show/{id} [get]
func getVendor(context *gin.Context) {
	id_, _ := strconv.ParseUint(context.Param("id"), 10, 32)
	id := uint(id_)

	code := http.StatusOK

	response := services.FindOneVendorById(id, *vendorRepository)

	if !response.Success {
		code = http.StatusBadRequest
	}

	context.JSON(code, response)
}

// @Summary Update a vendor
// @Description Updates a vendors
// @Tags vendors
// @Accept json
// @Produce json
// @Param id path string true "Vendor ID"
// @Param vendor body models.Vendor true "Vendor details"
// @Success 201 {object} dtos.Response
// @Failure 400 {object} dtos.Response
// @Router /vendors/update/{id} [put]
func updateVendor(context *gin.Context) {
	id_, _ := strconv.ParseUint(context.Param("id"), 10, 32)
	id := uint(id_)

	var vendor models.Vendor

	err := context.ShouldBindJSON(&vendor)

	if err != nil {
		context.JSON(http.StatusBadRequest, "Can't bind JSON to vendor")
		return
	}

	code := http.StatusOK

	response := services.UpdateVendorById(id, &vendor, *vendorRepository)

	if !response.Success {
		code = http.StatusBadRequest
	}

	context.JSON(code, response)
}

// @Summary Delete a vendor
// @Description Deletes a vendors
// @Tags vendors
// @Accept json
// @Produce json
// @Param id path string true "Vendor ID"
// @Success 201 {object} dtos.Response
// @Failure 400 {object} dtos.Response
// @Router /vendors/delete/{id} [delete]
func deleteVendor(context *gin.Context) {
	id_, _ := strconv.ParseUint(context.Param("id"), 10, 32)
	id := uint(id_)

	code := http.StatusOK

	response := services.DeleteOneVendorById(id, *vendorRepository)

	if !response.Success {
		code = http.StatusBadRequest
	}

	context.JSON(code, response)
}

// @Summary Delete vendors
// @Description Deletes vendors
// @Tags vendors
// @Accept json
// @Produce json
// @Param ids body dtos.MultiID true "Vendor IDs"
// @Success 201 {object} dtos.Response
// @Failure 400 {object} dtos.Response
// @Router /vendors/delete/ [post]
func deleteVendors(context *gin.Context) {
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

	response := services.DeleteVendorByIds(&multiID, *vendorRepository)

	if !response.Success {
		code = http.StatusBadRequest
	}

	context.JSON(code, response)
}

func SetupVendorsRoutes(vendorRepository_ *repository.VendorRepository, route *gin.Engine) {
	vendorRepository = vendorRepository_
	route.POST("/vendors/create", createVendor)
	route.GET("/vendors/", getVendors)
	route.GET("/vendors/show/:id", getVendor)
	route.PUT("/vendors/update/:id", updateVendor)
	route.DELETE("/vendors/delete/:id", deleteVendor)
	route.POST("/vendors/delete", deleteVendors)
}
