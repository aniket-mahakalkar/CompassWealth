package views

type AssetResponse struct {
	ID           uint    `json:"id"`
	Name         string  `json:"name"`
	Type         string  `json:"type"`
	NoOfUnits   float64 `json:"no_of_units"`
	AvgBuyPrice  float64 `json:"avg_buy_price"`
	CurrentPrice float64 `json:"current_price"`
	Charges      float64 `json:"charges"`
}

type AssetRequest struct {
	ID           uint    `json:"id"`
	Name         string  `json:"name"`
	Type         string  `json:"type"`
	NoOfUnits   float64 `json:"no_of_units"`
	AvgBuyPrice  float64 `json:"avg_buy_price"`
	CurrentPrice float64 `json:"current_price"`
	Charges      float64 `json:"charges"`
}

