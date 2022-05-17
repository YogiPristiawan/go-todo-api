package entities

import "time"

type GetUserByIdResponse struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	Gender    string    `json:"gender"`
	BirthDate time.Time `json:"birth_date"`
	CreatedAt uint      `json:"created_at"`
	UpdatedAt uint      `json:"updated_at"`
}

func MapGetUserByIdResponse(user *UserModel) *GetUserByIdResponse {
	userByIdResponse := &GetUserByIdResponse{
		ID:        user.ID,
		Username:  user.Username,
		Gender:    user.Gender,
		BirthDate: user.BirthDate,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return userByIdResponse
}
