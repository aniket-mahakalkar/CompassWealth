package model

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	UserName  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Location  string `json:"location"`
	IPAddress string `json:"ip_address"`
	Token     string `json:"token"`
}
