package auth

import (
	"net/http"
	"next-gen-job-hunting/api/user"
	"next-gen-job-hunting/common/exception"

	"github.com/gin-gonic/gin"
)

// AuthController handles authentication related operations
type AuthController struct {
	Service *AuthValidationService
}

// NewAuthController creates a new AuthController
func NewAuthController(validator *AuthValidationService) *AuthController {
	return &AuthController{
		Service: validator,
	}
}

// SignUp handles user registration
// @Summary Register a new user
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body user.User true "User"
// @Success 200 {object} exception.CommonException "Signup successful"
// @Failure 400 {object} exception.CommonException "Invalid input"
// @Failure 500 {object} exception.CommonException "Failed to generate token"
// @Router /auth/signup [post]
func (ctrl *AuthController) SignUp(c *gin.Context) {
	var user user.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, exception.NewCommonException("Invalid input", err.Error()))
		return
	}

	token, err := ctrl.Service.SignUp(&user, c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, exception.NewCommonException("Failed to generate token", err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Signup successful", "token": token.Token})
}

// SignIn handles user login
// @Summary Log in a user
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body user.User true "User (only email and password required)"
// @Success 200 {object} exception.CommonException "Signin successful"
// @Failure 400 {object} exception.CommonException "Invalid input"
// @Failure 500 {object} exception.CommonException "Failed to generate token"
// @Router /auth/signin [post]
func (ctrl *AuthController) SignIn(c *gin.Context) {
	var user user.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, exception.NewCommonException("Invalid input", err.Error()))
		return
	}

	token, err := ctrl.Service.SignIn(&user, c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, exception.NewCommonException("Failed to generate token", err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Signin successful", "token": token.Token})
}

// SignOut handles user logout
// @Summary Log out a user
// @Tags Auth
// @Produce json
// @Param auth_token header string true "Auth Token"
// @Success 200 {object} exception.CommonException "Signout successful"
// @Failure 500 {object} exception.CommonException "Failed to signout"
// @Router /auth/signout [post]
func (ctrl *AuthController) SignOut(c *gin.Context) {
	authToken := c.GetHeader("auth_token")

	err := ctrl.Service.SignOut(authToken, c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, exception.NewCommonException("Failed to signout", err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Signout successful"})
}
