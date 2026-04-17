package services

import (
	"compass-wealth/data"
	"compass-wealth/enums"
	"compass-wealth/model"
	"compass-wealth/utils"
	"compass-wealth/views"
	"encoding/json"
	"errors"
	"net/http"
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
		return errors.New("email is required")
	}

	if strings.TrimSpace(req.UserName) == "" {
		return errors.New("username is required")
	}

	if strings.TrimSpace(req.Password) == "" {
		return errors.New("password is required")
	}

	account.UserName = req.UserName
	account.Email = req.Email
	if req.Role != nil {
		account.Role = *req.Role
	} else {
		account.Role = enums.User
	}

	hashPas, err := utils.HashPassword(req.Password)
	if err != nil {
		return errors.New("error in hashing password")
	}
	account.Password = hashPas

	err = s.db.Where("email = ?", req.Email).First(&account).Error
	if err == nil {
		return errors.New("email already exists")
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("error checking for existing email")
	}

	if err := s.db.Create(&account).Error; err != nil {
		return errors.New("error in creating account")
	}

	return nil
}

func (s *AccountService) LoginAccount(req views.LoginAccount, ip string) (views.AccountResponse, error) {

	if strings.TrimSpace(req.Email) == "" {
		return views.AccountResponse{}, errors.New("email is required")
	}

	if strings.TrimSpace(req.Password) == "" {
		return views.AccountResponse{}, errors.New("password is required")
	}

	var account model.Account

	err := s.db.Where("email = ?", req.Email).First(&account).Error
	if err != nil {
		return views.AccountResponse{}, errors.New("email not found")
	}

	if !utils.CheckPasswordHash(req.Password, account.Password) {
		return views.AccountResponse{}, errors.New("invalid password or email")
	}

	token, err := utils.GenerateToken(account.ID, account.Email, account.Role)
	if err != nil {
		return views.AccountResponse{}, errors.New("error in generating token")
	}

	account.Token = token
	account.IPAddress = ip

	if err := s.db.Save(&account).Error; err != nil {
		return views.AccountResponse{}, errors.New("error in saving token")
	}

	return views.AccountResponse{
		ID:       account.ID,
		UserName: account.UserName,
		Email:    account.Email,
		Role:     account.Role,
		Token:    token,
	}, nil

}

func (s *AccountService) UpdateUsername(req views.AccountRequest) error {
	var account model.Account

	if req.ID == 0 {
		return errors.New("id is required")
	}

	if strings.TrimSpace(req.UserName) == "" {

		return errors.New("username is required")
	}

	if err := s.db.Where("id = ?", req.ID).First(&account).Error; err != nil {
		return errors.New("account not found")
	}

	account.UserName = req.UserName
	if err := s.db.Save(&account).Error; err != nil {
		return errors.New("error in updating username")
	}

	return nil
}

func GetLocationFromIP(ip string) (*string, error) {

	if ip == "::1" || ip == "127.0.0.1" {
		ip = "8.8.8.8"
	}

	url := "http://ip-api.com/json/" + ip

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var location string
	if err := json.NewDecoder(resp.Body).Decode(&location); err != nil {
		return nil, err
	}

	return &location, nil
}
