package model

import "gorm.io/gorm"

type Asset struct {
	gorm.Model
	Name         string
	Type         string
	NoOfUnits   float64
	AvgBuyPrice  float64
	CurrentPrice float64
	Charges      float64
	AccountID    uint
}