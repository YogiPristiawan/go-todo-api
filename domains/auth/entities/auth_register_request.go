package entities

type AuthRegisterRequest struct {
	Username  string  `json:"username" validate:"required,username"`
	Password  string  `json:"password" validate:"required,min=6"`
	Gender    *string `json:"gender" validate:"omitempty,eq=L|eq=P"`
	BirthDate *string `json:"birth_date" validate:"omitempty,datetime=2006-01-02"`
}
