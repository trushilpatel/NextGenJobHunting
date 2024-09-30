package user

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) CreateUser(user *User, c *gin.Context) error {
	if err := r.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) GetAllUser(c *gin.Context) ([]*User, error) {
	userContext, _ := c.Get("user")

	var users []*User
	if err := r.DB.Where("id = ?", userContext.(*User).ID.ID).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) GetUserByID(id uint, c *gin.Context) (*User, error) {
	var user User
	if err := r.DB.First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err // Record not found is not considered an error
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetUserByEmail(email string, c *gin.Context) (*User, error) {
	var user User
	result := r.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil // Record not found is not considered an error
		}
		return nil, result.Error
	}
	return &user, nil
}

func (r *UserRepository) UpdateUser(user *User, c *gin.Context) error {
	userContext, _ := c.Get("user")
	if err := r.DB.Where("Id = ?", userContext.(*User).ID.ID).Save(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) DeleteUser(id uint, c *gin.Context) error {
	userContext, _ := c.Get("user")
	if userContext.(*User).ID.ID != id {
		return nil
	}

	if err := r.DB.Delete(&User{}, id).Error; err != nil {
		return err
	}
	return nil
}
