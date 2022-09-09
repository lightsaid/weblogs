package service

type RegisterRequest struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	AckPassword string `json:"ack_password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Remember bool   `josn:"remember"`
}

type CreateUserRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	Avatar   string `json:"avatar"`
}

type SessionUser struct {
	UserID   int
	Username string
	Avatar   string
}

type CreatePostRequest struct {
	Title      string   `json:"title"`
	Content    string   `json:"content"`
	Attrs      []string `json:"attrs"`
	Categories []string `json:"categories"`
	Thumb      string   `json:"thumb"`
}
