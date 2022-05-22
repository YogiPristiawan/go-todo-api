package entities

import (
	"time"
)

type AuthRegisterRequest struct {
	Username  string    `json:"username" validate:"required,username"`
	Password  string    `json:"password" validate:"required,min=6"`
	Gender    string    `json:"gender" validate:"eq=L|eq=P"`
	BirthDate time.Time `json:"birth_date" validate:"datetime=2006-01-02"`
}
