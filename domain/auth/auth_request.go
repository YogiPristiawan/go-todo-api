package auth

import (
	"github.com/YogiPristiawan/go-todo-api/modules/alias"
)

type LoginRequest struct {
	Username string `json:"username" validate:"required,username"`
	Password string `json:"password" validate:"required,min=6"`
}

type RegisterRequest struct {
	Username  string           `json:"username" validate:"required,username"`
	Password  string           `json:"password" validate:"required,min=6"`
	Gender    alias.NullString `json:"gender" validate:"omitempty,eq=L|eq=P"`
	BirthDate alias.NullString `json:"birth_date" validate:"omitempty,datetime=2006-01-02"`
}
