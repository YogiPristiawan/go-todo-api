package user

type UserUseCase interface {
	Get() ([]*GetUsersResponse, error)
	FindById(id uint) (*GetUserByIdResponse, error)
}
