package user

import (
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) CreateUser(user *User) error {
	if err := r.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) GetAllUser() ([]*User, error) {
	var users []*User
	if err := r.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) GetUserByID(id uint) (*User, error) {
	var user User
	if err := r.DB.First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil // Record not found is not considered an error
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) UpdateUser(user *User) error {
	if err := r.DB.Save(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) DeleteUser(id uint) error {
	if err := r.DB.Delete(&User{}, id).Error; err != nil {
		return err
	}
	return nil
}
