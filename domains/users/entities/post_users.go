package entities

type PostUsers struct {
	Name      string `json:"name" validate:"required,alpha"`
	Gender    string `json:"gender" validate:"required,alpha"`
	BirthDate string `json:"birth_date" validate:"required,datetime=2006-01-02"`
	Weight    int    `json:"weight" validate:"required,numeric"`
	Height    int    `json:"height" validate:"required,numeric"`
}
