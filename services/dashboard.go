package services

import (
	"compass-wealth/data"
	"compass-wealth/middlewares"
	"compass-wealth/model"
	"compass-wealth/views"
	"errors"

	"gorm.io/gorm"
)

type DashboardService struct {
	db *gorm.DB
}

func NewDashboardService() *DashboardService {
	return &DashboardService{
		db: data.DB,
	}
}

func (s *DashboardService) GetDashboardData(user middlewares.AuthUser) (views.DashboardResponse, error) {

	if user.ID == 0 {
		return views.DashboardResponse{}, errors.New("invalid user")
	}

	var assets []views.DashboardAssetResponse
	var dashboard views.DashboardResponse

	if err := s.db.Model(&model.Asset{}).
		Select(`name,
	SUM(no_of_units * avg_buy_price) as total_investment,
	SUM(no_of_units * current_price) as total_current,
	SUM(no_of_units * (current_price - avg_buy_price)) as total_profit,
	SUM(charges) as total_charges
`).Where("account_id = ?", user.ID).Group("name").Scan(&assets).Error; err != nil {
		return views.DashboardResponse{}, errors.New("error fetching dashboard data")
	}

	dashboard.Assets = assets
	dashboard.TotalAssets = len(assets)
	dashboard.TotalInvestment = 0
	dashboard.TotalProfit = 0
	dashboard.TotalCharges = 0

	for _, asset := range assets {
		dashboard.TotalInvestment += asset.TotalInvestment
		dashboard.TotalProfit += asset.TotalProfit
		dashboard.TotalCharges += asset.TotalCharges
	}

	return dashboard, nil

}
