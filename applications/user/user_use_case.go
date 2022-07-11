package user

import (
	"github.com/YogiPristiawan/go-todo-api/domain/user"
)

type UserUseCase struct {
	UserRepository user.UserRepository
}

func NewUserUseCase(userRepository user.UserRepository) user.UserUseCase {
	return &UserUseCase{
		UserRepository: userRepository,
	}
}

func (u *UserUseCase) Get() ([]*user.GetUsersResponse, error) {
	users, err := u.UserRepository.Get()
	if err != nil {
		return nil, err
	}

	return user.MapGetUsersResponse(users), err
}

func (u *UserUseCase) FindById(userId uint) (*user.GetUserByIdResponse, error) {
	data, err := u.UserRepository.FindById(userId)
	if err != nil {
		return nil, err
	}
	return user.MapGetUserByIdResponse(data), err
}
