package entities

import (
	"time"
)

type GetUsersResponse struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	Gender    string    `json:"gender"`
	BirthDate time.Time `json:"birth_date"`
	CreatedAt uint      `json:"created_at"`
	UpdatedAt uint      `json:"updated_at"`
}

func MapGetUsersResponse(users []*UserModel) []*GetUsersResponse {
	var usersResponse []*GetUsersResponse

	for _, value := range users {
		usersResponse = append(usersResponse, &GetUsersResponse{
			ID:        value.ID,
			Username:  value.Username,
			Gender:    value.Gender,
			BirthDate: value.BirthDate,
			CreatedAt: value.CreatedAt,
			UpdatedAt: value.UpdatedAt,
		})
	}

	return usersResponse
}
