package repository

import (
	"github.com/YogiPristiawan/go-todo-api/domains/users"
	"github.com/YogiPristiawan/go-todo-api/domains/users/entities"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) users.UserRepository {
	return &userRepository{
		db,
	}
}

func (u *userRepository) GetAllUsers() []*entities.UserModel {
	var user []*entities.UserModel
	u.db.Find(&user)
	return user
}
