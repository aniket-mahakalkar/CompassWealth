package services

import (
	"compass-wealth/data"
	"compass-wealth/model"
	"compass-wealth/utils"
	"compass-wealth/views"
	"strings"

	"gorm.io/gorm"
)

type AccountService struct {
	db *gorm.DB
}

func NewAccountService() *AccountService {
	return &AccountService{
		db: data.DB,
	}
}

func (s *AccountService) CreateAccount(req views.AccountRequest) (error) {

	var account model.Account

	if strings.TrimSpace(req.Email) == "" {
		return  utils.ThrowError(nil, "email is required")
	}

	if strings.TrimSpace(req.UserName) == "" {
		return  utils.ThrowError(nil, "username is required")
	}

	if strings.TrimSpace(req.Password) == "" {
		return  utils.ThrowError(nil, "password is required")
	}

	account.UserName = req.UserName
	account.Email = req.Email
	account.Password = req.Password


	if err := s.db.Where("email = ?", req.Email).First(&account).Error; err != nil {
		return  utils.ThrowError(err, "email already exists")
	}

	return nil
}
