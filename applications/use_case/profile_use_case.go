package use_case

import (
	"github.com/YogiPristiawan/go-todo-api/domains/profile"
	"github.com/YogiPristiawan/go-todo-api/domains/profile/entities"
	"github.com/YogiPristiawan/go-todo-api/domains/users"
)

type ProfileUseCase struct {
	userRepository users.UserRepository
}

func NewProfileUseCase(userRepository users.UserRepository) profile.ProfileUseCase {
	return &ProfileUseCase{
		userRepository: userRepository,
	}
}

func (p *ProfileUseCase) GetProfile(userId uint) (*entities.GetProfileResponse, error) {
	user, err := p.userRepository.GetUserById(userId)
	if err != nil {
		return nil, err
	}

	return entities.MapGetProfileResponse(user), nil
}
