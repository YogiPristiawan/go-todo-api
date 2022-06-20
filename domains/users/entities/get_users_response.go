package entities

import "fmt"

type GetUsersResponse struct {
	ID        uint    `json:"id"`
	Username  string  `json:"username"`
	Gender    *string `json:"gender"`
	BirthDate *string `json:"birth_date"`
	CreatedAt int64   `json:"created_at"`
	UpdatedAt int64   `json:"updated_at"`
}

func MapGetUsersResponse(users []*UserModel) []*GetUsersResponse {
	var usersResponse []*GetUsersResponse

	if len(users) == 0 {
		return make([]*GetUsersResponse, 0)
	}

	for _, value := range users {
		fmt.Println("masuk ke looping")
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
