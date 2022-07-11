package auth

type AuthUseCase interface {
	Login(*LoginRequest) (*LoginResponse, error)
	Register(*RegisterRequest) (*RegisterResponse, error)
}
