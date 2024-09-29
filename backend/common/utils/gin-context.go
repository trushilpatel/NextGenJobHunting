package utils

// import (
// 	"errors"
// 	"os/user"

// 	"github.com/gin-gonic/gin"
// )

// // getUser retrieves the user from the context and performs type assertion
// func (c *gin.Context) getUser() (*user.User, error) {
// 	userContext, exists := c.Get("user")
// 	if !exists {
// 		return nil, errors.New("user not found in context")
// 	}
// 	user, ok := userContext.(*user.User)
// 	if !ok {
// 		return nil, errors.New("user type assertion failed")
// 	}
// 	return user, nil
// }

// // getUser retrieves the user from the context and performs type assertion
// func (c *gin.Context) setUser(user *user.User) (*user.User, error) {
// 	c.Set("user", user)
// 	return user, nil
// }
