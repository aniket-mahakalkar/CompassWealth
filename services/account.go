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

func (s *AccountService) CreateAccount(req views.AccountRequest) error {

	var account model.Account

	if strings.TrimSpace(req.Email) == "" {
		return utils.ThrowError(nil, "email is required")
	}

	if strings.TrimSpace(req.UserName) == "" {
		return utils.ThrowError(nil, "username is required")
	}

	if strings.TrimSpace(req.Password) == "" {
		return utils.ThrowError(nil, "password is required")
	}

	account.UserName = req.UserName
	account.Email = req.Email
	account.Password = req.Password

	if err := s.db.Where("email = ?", req.Email).First(&account).Error; err != nil {
		return utils.ThrowError(err, "email already exists")
	}

	return nil
}

func (s *AccountService) LoginAccount(req views.LoginAccount) (views.AccountResponse, error) {

	if strings.TrimSpace(req.UserName) == "" {
		return views.AccountResponse{}, utils.ThrowError(nil, "username is required")
	}

	if strings.TrimSpace(req.Password) == "" {
		return views.AccountResponse{}, utils.ThrowError(nil, "password is required")
	}

	var account model.Account

	if err := s.db.Where("username = ?", req.UserName).First(&account).Error; err != nil {
		return views.AccountResponse{}, utils.ThrowError(err, "username not found")
	}

	if !utils.CheckPasswordHash(req.Password, account.Password) {
		return views.AccountResponse{}, utils.ThrowError(nil, "invalid password")
	}

	return views.AccountResponse{
		ID:       account.ID,
		UserName: account.UserName,
		Email:    account.Email,
	}, nil

}
