package profile

type ProfileUseCase interface {
	FindByUserId(userId uint) (*FindProfileResponse, error)
}
