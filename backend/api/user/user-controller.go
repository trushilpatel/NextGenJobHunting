package user

import (
	"log"
	"net/http"
	"next-gen-job-hunting/common/exception"
	"strconv"

	"github.com/gin-gonic/gin"
)

// UserController handles user-related requests
type UserController struct {
	Service *UserValidationService
}

// NewUserController creates a new UserController
func NewUserController(service *UserValidationService) *UserController {
	return &UserController{Service: service}
}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user with the input payload
// @Tags users
// @Accept json
// @Produce json
// @Param user body User true "User"
// @Success 201 {object} User
// @Failure 400 {object} exception.CommonException "Invalid request payload"
// @Failure 500 {object} exception.CommonException "Service not initialized" "User not initialized" "Unable to create user"
// @Router /users [post]
func (c *UserController) CreateUser(ctx *gin.Context) {
	var user *User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, exception.NewCommonException("Invalid request payload", err.Error()))
		return
	}

	if c.Service == nil {
		ctx.JSON(http.StatusInternalServerError, exception.NewCommonException("Service not initialized", ""))
		return
	}

	if user == nil {
		ctx.JSON(http.StatusInternalServerError, exception.NewCommonException("User not initialized", ""))
		return
	} else {
		log.Println(user)
	}

	if err := c.Service.CreateUser(user, ctx); err != nil {
		ctx.JSON(http.StatusInternalServerError, exception.NewCommonException("Unable to create user", err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, user)
}

// GetAllUser godoc
// @Summary Get all users
// @Description Get all users
// @Tags users
// @Produce json
// @Success 200 {array} User
// @Failure 400 {object} exception.CommonException "Error occurred while retrieving users"
// @Router /users [get]
func (c *UserController) GetAllUser(ctx *gin.Context) {
	users, err := c.Service.GetAllUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, exception.NewCommonException("Error occurred while retrieving users", err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, users)
}

// GetUser godoc
// @Summary Get a user by ID
// @Description Get a user by ID
// @Tags users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} User
// @Failure 400 {object} exception.CommonException "Invalid user ID"
// @Failure 404 {object} exception.CommonException "User not found"
// @Router /users/{id} [get]
func (c *UserController) GetUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, exception.NewCommonException("Invalid user ID", err.Error()))
		return
	}

	user, err := c.Service.GetUserByID(uint(id), ctx)
	if err != nil {
		ctx.JSON(http.StatusNotFound, exception.NewCommonException("User not found", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, user)
}

// UpdateUser godoc
// @Summary Update a user by ID
// @Description Update a user by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body User true "User"
// @Success 200 {object} User
// @Failure 400 {object} exception.CommonException "Invalid user ID" "Invalid request payload"
// @Failure 500 {object} exception.CommonException "Unable to update user"
// @Router /users/{id} [put]
func (c *UserController) UpdateUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, exception.NewCommonException("Invalid user ID", err.Error()))
		return
	}

	var user User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, exception.NewCommonException("Invalid request payload", err.Error()))
		return
	}

	user.ID.ID = uint(id)
	if err := c.Service.UpdateUser(&user, ctx); err != nil {
		ctx.JSON(http.StatusInternalServerError, exception.NewCommonException("Unable to update user", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, user)
}

// DeleteUser godoc
// @Summary Delete a user by ID
// @Description Delete a user by ID
// @Tags users
// @Param id path int true "User ID"
// @Success 204
// @Failure 400 {object} exception.CommonException "Invalid user ID"
// @Failure 500 {object} exception.CommonException "Unable to delete user"
// @Router /users/{id} [delete]
func (c *UserController) DeleteUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, exception.NewCommonException("Invalid user ID", err.Error()))
		return
	}

	if err := c.Service.DeleteUser(uint(id), ctx); err != nil {
		ctx.JSON(http.StatusInternalServerError, exception.NewCommonException("Unable to delete user", err.Error()))
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
