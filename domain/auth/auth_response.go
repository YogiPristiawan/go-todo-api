package auth

type RegisterResponse struct {
	AccessToken string `json:"access_token"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}
