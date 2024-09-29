package user

import (
	"github.com/gin-gonic/gin"
)

type IUserService interface {
	NewUserService(repo *UserRepository) *UserService
	CreateUser(user *User) error
	GetAllUser() ([]*User, error)
	GetUser(id uint) (*User, error)
	UpdateUser(user *User) error
	DeleteUser(id uint) error
}

func RegisterUserRoutes(routerGroup *gin.RouterGroup, controller *UserController) {
	userGroup := routerGroup.Group("/user")
	{
		//userGroup.GET("", controller.GetAllUser)
		userGroup.GET("/:id", controller.GetUser)
		userGroup.POST("", controller.CreateUser)
		userGroup.PUT("", controller.UpdateUser)
		userGroup.DELETE("", controller.DeleteUser)
	}
}
