package dto

import "go_todo_api/src/shared/entities"

// ProfileRequest provides data struct to
// show user profile
type ProfileRequest struct {
	entities.RequestMetaData
}

// ProfileResponse provides data struct
// to show user profile response
type ProfileResponse struct {
	Id        int64           `json:"id"`
	Username  string          `json:"username"`
	Gender    entities.String `json:"gender"`
	BirthDate entities.String `json:"birth_date"`
}
