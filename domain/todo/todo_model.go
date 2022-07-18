package todo

type (
	Todo_ID         uint
	Todo_UserId     uint
	Todo_Todo       string
	Todo_Date       string
	Todo_IsFinished bool
	Todo_CreatedAt  int64
	Todo_UpdatedAt  int64
)

type TodoModel struct {
	ID         uint
	UserId     uint
	Todo       string
	Date       string
	IsFinished bool
	CreatedAt  int64 `gorm:"autoCreateTime"`
	UpdatedAt  int64 `gorm:"autoUpdateTime"`
}

func (TodoModel) TableName() string {
	return "todos"
}
