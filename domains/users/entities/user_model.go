package entities

type UserModel struct {
	ID        uint
	Username  string
	Password  string
	Gender    *string
	BirthDate *string
	CreatedAt int64 `gorm:"autoCreateTime"`
	UpdatedAt int64 `gorm:"autoUpdateTime"`
}

func (UserModel) TableName() string {
	return "users"
}
