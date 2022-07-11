package user

type UserRepository interface {
	Store(*UserModel) (*UserModel, error)
	Get() ([]*UserModel, error)
	FindById(id uint) (*UserModel, error)
	FindByUsername(username string) (*UserModel, error)

	VerifyAvailableUsername(username string) error
}
