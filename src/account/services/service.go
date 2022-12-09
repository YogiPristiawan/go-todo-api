package services

import (
	"go_todo_api/src/account/dto"
	"go_todo_api/src/account/models"
	"go_todo_api/src/shared/databases"
	"go_todo_api/src/shared/jwt"

	"golang.org/x/crypto/bcrypt"
)

// mockError is a struct to mock an error, for testing purpose
type mockError struct {
	code int
}

func (h *mockError) Error() string {
	return ""
}

// this variables store helper function
// for easier testing
var wrapDBErr = databases.WrapError
var comparePassword = func(hashedPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
var hashPassword = func(password string) string {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(hashed)
}
var generateAccessToken = jwt.GenerateAccessToken

func mapGetProfileToResponse(res *dto.ProfileResponse, models *models.Profile) {
	res.Id = models.Id
	res.Username = models.Username
	res.Gender = models.Gender
	res.BirthDate = models.BirthDate
}
