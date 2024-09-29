package auth

import (
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(router *gin.RouterGroup, authController *AuthController) {
	authRoutes := router.Group("/auth")
	{
		authRoutes.GET("/signup", authController.SignUp)
		authRoutes.POST("/signin", authController.SignIn)
		authRoutes.POST("/signout", authController.SignOut)
	}
}
