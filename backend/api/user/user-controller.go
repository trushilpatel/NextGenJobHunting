package user

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	Service *UserValidationService
}

func NewUserController(service *UserValidationService) *UserController {
	return &UserController{Service: service}
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	var user *User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if c.Service == nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Service not initialized"})
		return
	}

	if user == nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "User not initialized"})
		return
	} else {
		log.Println(user)
	}

	if err := c.Service.CreateUser(user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create user"})
		return
	}

	ctx.JSON(http.StatusCreated, user)
}

func (c *UserController) GetAllUser(ctx *gin.Context) {
	users, err := c.Service.GetAllUser()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error occured while retrieving users"})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (c *UserController) GetUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := c.Service.GetUserByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (c *UserController) UpdateUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var user User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.ID.ID = uint(id)
	if err := c.Service.UpdateUser(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (c *UserController) DeleteUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	if err := c.Service.DeleteUser(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
