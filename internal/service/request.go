package service

type LoginWithRegisterRequest struct {
	FormType    string `json:"formType"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	AckPassword string `json:"ack_password"`
	Remember    bool   `josn:"remember"`
}

type CreateUserRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	Avatar   string `json:"avatar"`
}