package views

type DashboardAssetResponse struct {
	Name            string  `json:"name"            db:"name"`
	NoOfUnits       float64 `json:"no_of_units"      db:"total_unit"`
	TotalInvestment float64 `json:"total_investment" db:"total_investment"`
	TotalCurrent    float64 `json:"total_current"    db:"total_current"`
	TotalProfit     float64 `json:"total_profit"     db:"total_profit"`
	TotalCharges    float64 `json:"total_charges"    db:"total_charges"`
}

type DashboardResponse struct {
	TotalAssets     int                `json:"total_assets"`
	TotalInvestment float64            `json:"total_investment"`
	TotalProfit     float64            `json:"total_profit"`
	TotalCharges    float64            `json:"total_charges"`
	Assets          []DashboardAssetResponse `json:"assets"`
}
