package auth

import (
	"net/http"
	"next-gen-job-hunting/api/user"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	Service *AuthValidationService
}

func NewAuthController(validator *AuthValidationService) *AuthController {
	return &AuthController{
		Service: validator,
	}
}

func (ctrl *AuthController) SignUp(c *gin.Context) {
	var user user.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := ctrl.Service.SignUp(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token", "message": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Signup successful", "token": token.Token})
}

func (ctrl *AuthController) SignIn(c *gin.Context) {
	var user user.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := ctrl.Service.SignIn(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Signin successful", "token": token.Token})
}

func (ctrl *AuthController) SignOut(c *gin.Context) {
	var token string
	if err := c.ShouldBindJSON(&token); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Signout successful"})
}
