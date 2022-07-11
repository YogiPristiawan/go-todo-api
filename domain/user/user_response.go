package user

import (
	"github.com/YogiPristiawan/go-todo-api/modules/alias"
)

type GetUsersResponse struct {
	ID        uint             `json:"id"`
	Username  string           `json:"username"`
	Gender    alias.NullString `json:"gender"`
	BirthDate alias.NullString `json:"birth_date"`
	CreatedAt int64            `json:"created_at"`
	UpdatedAt int64            `json:"updated_at"`
}

func MapGetUsersResponse(users []*UserModel) []*GetUsersResponse {
	var usersResponse []*GetUsersResponse

	if len(users) == 0 {
		return make([]*GetUsersResponse, 0)
	}

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

type GetUserByIdResponse struct {
	ID        uint             `json:"id"`
	Username  string           `json:"username"`
	Gender    alias.NullString `json:"gender"`
	BirthDate alias.NullString `json:"birth_date"`
	CreatedAt int64            `json:"created_at"`
	UpdatedAt int64            `json:"updated_at"`
}

func MapGetUserByIdResponse(user *UserModel) *GetUserByIdResponse {
	userById := &GetUserByIdResponse{
		ID:        user.ID,
		Username:  user.Username,
		Gender:    user.Gender,
		BirthDate: user.BirthDate,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return userById
}
