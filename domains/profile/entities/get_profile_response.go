package entities

import "github.com/YogiPristiawan/go-todo-api/domains/users/entities"

type GetProfileResponse struct {
	ID        uint    `json:"id"`
	Username  string  `json:"username"`
	Gender    *string `json:"gender"`
	BirthDate *string `json:"birth_date"`
}

func MapGetProfileResponse(user *entities.UserModel) *GetProfileResponse {
	return &GetProfileResponse{
		ID:        user.ID,
		Username:  user.Username,
		Gender:    user.Gender,
		BirthDate: user.BirthDate,
	}
}
