package routes

import (
	"net/http"
	"strconv"
	"tspo_final/internal/dtos"
	"tspo_final/internal/models"
	"tspo_final/internal/repository"
	services "tspo_final/internal/services/user"

	"github.com/gin-gonic/gin"
)

var userRepository *repository.UserRepository

// @Summary Create user
// @Description Creates a users
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.User true "User details"
// @Success 201 {object} dtos.Response
// @Failure 400 {object} dtos.Response
// @Router /users/create [post]
func createUser(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, "Can't bind JSON to user")
		return
	}

	code := http.StatusOK

	response := services.CreateUser(&user, *userRepository)

	if !response.Success {
		code = http.StatusBadRequest
	}

	context.JSON(code, response)
}

// @Summary Get all users
// @Description Returns list of users
// @Tags users
// @Produce json
// @Success 200 {object} dtos.Response
// @Failure 400 {object} dtos.Response
// @Router /users/ [get]
func getUsers(context *gin.Context) {
	code := http.StatusOK

	response := services.FindAllUsers(*userRepository)

	if !response.Success {
		code = http.StatusBadRequest
	}

	context.JSON(code, response)
}

// @Summary Get one user
// @Description Returns one user
// @Tags users
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} dtos.Response
// @Failure 400 {object} dtos.Response
// @Router /users/show/{id} [get]
func getUser(context *gin.Context) {
	id_, _ := strconv.ParseUint(context.Param("id"), 10, 32)
	id := uint(id_)

	code := http.StatusOK

	response := services.FindOneUserById(id, *userRepository)

	if !response.Success {
		code = http.StatusBadRequest
	}

	context.JSON(code, response)
}

// @Summary Update a user
// @Description Updates a users
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param user body models.User true "User details"
// @Success 201 {object} dtos.Response
// @Failure 400 {object} dtos.Response
// @Router /users/update/{id} [put]
func updateUser(context *gin.Context) {
	id_, _ := strconv.ParseUint(context.Param("id"), 10, 32)
	id := uint(id_)

	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, "Can't bind JSON to user")
		return
	}

	code := http.StatusOK

	response := services.UpdateUserById(id, &user, *userRepository)

	if !response.Success {
		code = http.StatusBadRequest
	}

	context.JSON(code, response)
}

// @Summary Delete a user
// @Description Deletes a users
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 201 {object} dtos.Response
// @Failure 400 {object} dtos.Response
// @Router /users/delete/{id} [delete]
func deleteUser(context *gin.Context) {
	id_, _ := strconv.ParseUint(context.Param("id"), 10, 32)
	id := uint(id_)

	code := http.StatusOK

	response := services.DeleteOneUserById(id, *userRepository)

	if !response.Success {
		code = http.StatusBadRequest
	}

	context.JSON(code, response)
}

// @Summary Delete users
// @Description Deletes users
// @Tags users
// @Accept json
// @Produce json
// @Param ids body dtos.MultiID true "User IDs"
// @Success 201 {object} dtos.Response
// @Failure 400 {object} dtos.Response
// @Router /users/delete/ [post]
func deleteUsers(context *gin.Context) {
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

	response := services.DeleteUserByIds(&multiID, *userRepository)

	if !response.Success {
		code = http.StatusBadRequest
	}

	context.JSON(code, response)
}

func SetupUsersRoutes(userRepository_ *repository.UserRepository, route *gin.Engine) {
	userRepository = userRepository_
	route.POST("/users/create", createUser)
	route.GET("/users/", getUsers)
	route.GET("/users/show/:id", getUser)
	route.PUT("/users/update/:id", updateUser)
	route.DELETE("/users/delete/:id", deleteUser)
	route.POST("/users/delete", deleteUsers)
}
