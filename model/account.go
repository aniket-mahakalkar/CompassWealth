package model

import (
	"compass-wealth/enums"

	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	UserName  string
	Password  string
	Email     string `gorm:"uniqueIndex"`
	Role      enums.Roles
	Location  string
	IPAddress string
	Token     string
	Assets    []Asset  `gorm:"foreignKey:AccountID"`
}
