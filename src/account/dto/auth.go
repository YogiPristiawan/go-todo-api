package dto

import "go_todo_api/src/shared/entities"

// LoginRequest provides data struct of login request
type LoginRequest struct {
	Username string `json:"username" validate:"requried,username"`
	Password string `json:"password" validate:"requred,min=6"`
}

// LoginResponse provides data struct of login response
type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

// RegisterRequest provides data struct of register request
type RegisterRequest struct {
	Username  string          `json:"username" validate:"required,username"`
	Password  string          `json:"password" validate:"required,min=6"`
	Gender    entities.String `json:"gender" validate:"omitempty,eq=L|eq=P"`
	BirthDate entities.Date   `json:"birth_date" validate:"omitempty,datetime=2006-01-02"`
}

// RegisterResponse provides data struct of register response
type RegisterResponse struct {
	AccessToken string `json:"access_token"`
}
