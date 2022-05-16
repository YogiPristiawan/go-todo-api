package entities

import "time"

type UserModel struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Gender    string    `json:"gender"`
	BirthDate time.Time `json:"birth_date"`
	CreatedAt uint      `json:"created_at"`
	UpdatedAt uint      `json:"updated_at"`
}

func (UserModel) TableName() string {
	return "users"
}
