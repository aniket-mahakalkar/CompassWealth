package views

type AccountRequest struct {
	ID       uint   `json:"id"`
	UserName string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type LoginAccount struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AccountResponse struct {
	ID       uint   `json:"id"`
	UserName string `json:"username"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}
