package entities

type GetUsersResponse struct {
	Name      string `json:"name"`
	Gender    string `json:"gender"`
	BirthDate string `json:"birth_date"`
	Height    int    `json:"height"`
	Weight    int    `json:"weight"`
}
