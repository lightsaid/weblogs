package service

type LoginWithRegisterResponse struct {
	IfAdmin int `json:"if_admin"`
}

type JSONResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}
