package auth

import "github.com/YogiPristiawan/go-todo-api/domains/auth/entities"

type AuthUseCase interface {
	Login(*entities.AuthLoginRequest) (*entities.AuthLoginResponse, error)
}
