package models

import "go_todo_api/src/shared/entities"

// Account provides data struct
// of account data to interact with database
type Account struct {
	Id        int64
	Username  string
	Password  string
	Gender    entities.String
	BirthDate entities.Date
	CreatedAt int64
	UpdatedAt int64
}

// Profile struct provides data struct
// of profile data to interact with database
type Profile struct {
	Id        int64
	Username  string
	Gender    entities.String
	BirthDate entities.String
}
