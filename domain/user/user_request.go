package user

import (
	"github.com/YogiPristiawan/go-todo-api/modules/alias"
)

type StoreUserRequest struct {
	Name      string           `json:"name" validate:"required,alpha"`
	Gender    alias.NullString `json:"gender" validate:"alpha"`
	BirthDate alias.NullString `json:"birth_date" validate:"datetime=2006-01-02"`
	Weight    int              `json:"weight" validate:"required,numeric"`
	Height    int              `json:"height" validate:"required,numeric"`
}
