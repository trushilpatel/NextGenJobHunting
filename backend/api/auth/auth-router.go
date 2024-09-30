package auth

import (
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(router *gin.RouterGroup, authController *AuthController) {
	authRoutes := router.Group("/auth")
	{
		authRoutes.POST("/signup", authController.SignUp)
		authRoutes.POST("/signin", authController.SignIn)
		authRoutes.GET("/signout", authController.SignOut)
	}
}
