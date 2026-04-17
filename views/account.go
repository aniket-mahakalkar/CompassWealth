package views

import "compass-wealth/enums"

type AccountRequest struct {
	ID       uint         `json:"id"`
	UserName string       `json:"username"`
	Password string       `json:"password"`
	Email    string       `json:"email"`
	Role     *enums.Roles `json:"role,omitempty"`
}

type LoginAccount struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AccountResponse struct {
	ID       uint        `json:"id"`
	UserName string      `json:"username"`
	Email    string      `json:"email"`
	Role     enums.Roles `json:"role"`
	Token    string      `json:"token,omitempty"`
}
