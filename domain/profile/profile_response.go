package profile

import (
	"go_todo_api/domain/user"
	"go_todo_api/modules/alias"
)

type FindProfileResponse struct {
	ID        uint             `json:"id"`
	Username  string           `json:"username"`
	Gender    alias.NullString `json:"gender"`
	BirthDate alias.NullString `json:"birth_date"`
}

func MapFindProfileResponse(user *user.UserModel) *FindProfileResponse {
	return &FindProfileResponse{
		ID:        user.ID,
		Username:  user.Username,
		Gender:    user.Gender,
		BirthDate: user.BirthDate,
	}
}
