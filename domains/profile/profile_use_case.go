package profile

import "github.com/YogiPristiawan/go-todo-api/domains/profile/entities"

type ProfileUseCase interface {
	GetProfile(userId uint) (*entities.GetProfileResponse, error)
}
