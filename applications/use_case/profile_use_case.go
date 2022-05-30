package use_case

import (
	"github.com/YogiPristiawan/go-todo-api/domains/profile/entities"
	"github.com/YogiPristiawan/go-todo-api/domains/users"
)

type ProfileUseCase struct {
	UserRepository users.UserRepository
}

func (p *ProfileUseCase) GetProfile(userId uint) (*entities.GetProfileResponse, error) {
	user, err := p.UserRepository.GetUserById(userId)
	if err != nil {
		return nil, err
	}

	return entities.MapGetProfileResponse(user), nil
}
