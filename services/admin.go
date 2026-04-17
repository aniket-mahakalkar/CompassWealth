package services

import (
	"compass-wealth/data"
	"compass-wealth/model"
	"compass-wealth/views"
	"errors"

	"gorm.io/gorm"
)

type AdminService struct {
	db *gorm.DB
}

func NewAdminService() *AdminService {
	return &AdminService{
		db: data.DB,
	}
}

func (s AdminService) ListAllUsers() ([]views.AccountResponse, error) {

	var accounts []model.Account

	if err := s.db.Find(&accounts).Error; err != nil {
		return nil, errors.New("error fetching users")
	}

	var responses []views.AccountResponse
	for _, acc := range accounts {
		responses = append(responses, views.AccountResponse{
			ID:       acc.ID,
			UserName: acc.UserName,
			Email:    acc.Email,
			Role:     acc.Role,
		})
	}
	return responses, nil

}

func (s *AdminService) UpdateUserRole(req views.EditUserRole) error {
	var account model.Account
	if err := s.db.First(&account, uint(req.ID)).Error; err != nil {
		return errors.New("user not found")
	}

	account.Role = req.Role
	if err := s.db.Save(&account).Error; err != nil {
		return errors.New("error updating role")
	}

	return nil
}
