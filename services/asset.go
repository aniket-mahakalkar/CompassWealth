package services

import (
	"compass-wealth/data"
	"compass-wealth/middlewares"
	"compass-wealth/model"
	"compass-wealth/utils"
	"compass-wealth/views"
	"errors"
	"strings"

	"gorm.io/gorm"
)

type AssetService struct {
	db *gorm.DB
}

func NewAssetService() *AssetService {
	return &AssetService{
		db: data.DB,
	}
}

func (s *AssetService) GetAssets(accountID uint) ([]views.AssetResponse, error) {
	var assets []model.Asset
	if err := s.db.Where("account_id = ?", accountID).Find(&assets).Error; err != nil {
		return nil, errors.New("error fetching assets")
	}

	response := make([]views.AssetResponse, 0, len(assets))
	for _, asset := range assets {
		response = append(response, views.AssetResponse{
			ID:           asset.ID,
			Name:         asset.Name,
			Type:         asset.Type,
			NoOfUnits:   asset.NoOfUnits,
			AvgBuyPrice:  asset.AvgBuyPrice,
			CurrentPrice: asset.CurrentPrice,
			Charges:      asset.Charges,
		})
	}

	return response, nil
}

func (s *AssetService) CreateAsset(user middlewares.AuthUser, req views.AssetRequest) error {

	var asset model.Asset

	if user.ID == 0 {
		return utils.ThrowError(nil, "invalid user")
	}

	if strings.TrimSpace(req.Name) == "" {
		return utils.ThrowError(nil, "name is required")
	}

	if strings.TrimSpace(req.Type) == "" {
		return utils.ThrowError(nil, "type is required")
	}
	if req.NoOfUnits == 0 {
		return utils.ThrowError(nil, "no_of_stocks is required")
	}
	if req.AvgBuyPrice == 0 {
		return utils.ThrowError(nil, "avg_buy_price is required")
	}

	if req.CurrentPrice == 0 {
		return utils.ThrowError(nil, "current_price is required")
	}

	asset.Name = req.Name
	asset.Type = req.Type
	asset.NoOfUnits = req.NoOfUnits
	asset.AvgBuyPrice = req.AvgBuyPrice
	asset.CurrentPrice = req.CurrentPrice
	asset.Charges = req.Charges
	asset.AccountID = user.ID

	if err := s.db.Create(&asset).Error; err != nil {
		return utils.ThrowError(err, "error in creating asset")
	}

	return nil
}

func (s *AssetService) UpdateAsset(user middlewares.AuthUser, req views.AssetRequest) error {

	var asset model.Asset

	if user.ID == 0 {
		return utils.ThrowError(nil, "invalid user")
	}

	if req.ID == 0 {
		return utils.ThrowError(nil, "invalid asset id")
	}

	if strings.TrimSpace(req.Name) != "" {
		asset.Name = req.Name
	}

	if strings.TrimSpace(req.Type) != "" {
		asset.Type = req.Type
	}

	if req.NoOfUnits != 0 {
		asset.NoOfUnits = req.NoOfUnits
	}

	if req.AvgBuyPrice != 0 {
		asset.AvgBuyPrice = req.AvgBuyPrice
	}

	if req.CurrentPrice != 0 {
		asset.CurrentPrice = req.CurrentPrice
	}

	if req.Charges != 0 {
		asset.Charges = req.Charges
	}

	if err := s.db.Where("id = ? AND account_id = ?", req.ID, user.ID).First(&asset).Error; err != nil {
		return utils.ThrowError(err, "error fetching asset")
	}

	if asset.AccountID != user.ID {
		return utils.ThrowError(nil, "unauthorized to update this asset")
	}

	if err := s.db.Save(&asset).Error; err != nil {
		return utils.ThrowError(err, "error in updating asset")
	}

	return  nil
}


func (s *AssetService) DeleteAsset(user middlewares.AuthUser, req views.AssetRequest) error {

	var asset model.Asset

	if user.ID == 0 {
		return utils.ThrowError(nil, "invalid user")
	}

	if req.ID == 0 {
		return utils.ThrowError(nil, "invalid asset id")
	}

	if err := s.db.Where("id = ? AND account_id = ?", req.ID, user.ID).First(&asset).Error; err != nil {
		return utils.ThrowError(err, "error fetching asset")
	}

	if asset.AccountID != user.ID {
		return utils.ThrowError(nil, "unauthorized to delete this asset")
	}

	if err := s.db.Delete(&asset).Error; err != nil {
		return utils.ThrowError(err, "error in deleting asset")
	}

	return  nil
}