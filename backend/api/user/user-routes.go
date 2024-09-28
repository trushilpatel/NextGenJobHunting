package user

import (
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(routerGroup *gin.RouterGroup, controller *UserController) {
	userGroup := routerGroup.Group("/user")
	{
		userGroup.GET("/:id", controller.GetUser)
		userGroup.GET("", controller.GetAllUser)
		userGroup.POST("", controller.CreateUser)
		userGroup.PUT("", controller.UpdateUser)
		userGroup.DELETE("", controller.DeleteUser)
	}
}
