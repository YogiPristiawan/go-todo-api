package profile

import (
	"go_todo_api/domain/profile"
	"go_todo_api/domain/user"
)

type ProfileUseCase struct {
	UserRepository user.UserRepository
}

func NewProfileUseCase(userRepository user.UserRepository) profile.ProfileUseCase {
	return &ProfileUseCase{
		UserRepository: userRepository,
	}
}

func (p *ProfileUseCase) FindByUserId(userId uint) (*profile.FindProfileResponse, error) {
	user, err := p.UserRepository.FindById(userId)
	if err != nil {
		return nil, err
	}

	return profile.MapFindProfileResponse(user), nil
}
