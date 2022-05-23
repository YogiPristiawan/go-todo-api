package entities

type GetUserByIdResponse struct {
	ID        uint    `json:"id"`
	Username  string  `json:"username"`
	Gender    *string `json:"gender"`
	BirthDate *string `json:"birth_date"`
	CreatedAt int64   `json:"created_at"`
	UpdatedAt int64   `json:"updated_at"`
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
