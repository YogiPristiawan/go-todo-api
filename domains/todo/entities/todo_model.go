package entities

type TodoModel struct {
	ID         uint
	UserId     uint
	Todo       string
	Date       string
	IsFinished *bool
	CreatedAt  int64 `gorm:"autoCreateTime"`
	UpdatedAt  int64 `gorm:"autoUpdateTime"`
}

func (TodoModel) TableName() string {
	return "todos"
}
