package app

type UserCreateRequest struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	Description string `json:"description"`
}

type UserCreateResponse struct {
}
