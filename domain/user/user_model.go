package user

import (
	"github.com/YogiPristiawan/go-todo-api/modules/alias"
)

type UserModel struct {
	ID        uint
	Username  string
	Password  string
	Gender    alias.NullString
	BirthDate alias.NullString
	CreatedAt int64 `gorm:"autoCreateTime"`
	UpdatedAt int64 `gorm:"autoUpdateTime"`
}

func (UserModel) TableName() string {
	return "users"
}
