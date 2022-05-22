package entities

import "time"

type UserModel struct {
	ID        uint
	Username  string
	Password  string
	Gender    string
	BirthDate time.Time
	CreatedAt uint
	UpdatedAt uint
}

func (UserModel) TableName() string {
	return "users"
}
