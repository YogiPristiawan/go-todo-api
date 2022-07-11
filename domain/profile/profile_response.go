package profile

import (
	"github.com/YogiPristiawan/go-todo-api/domain/user"
	"github.com/YogiPristiawan/go-todo-api/modules/alias"
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
