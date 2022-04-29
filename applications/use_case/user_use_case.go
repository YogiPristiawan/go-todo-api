package use_case

import "github.com/YogiPristiawan/go-todo-api/infrastructures/repository"

type UserUseCase struct {
	userRepository *repository.UserRepository
}

func (u *UserUseCase) GetAllUsers() []map[string]interface{} {
	return u.userRepository.GetAllUsers()
}
