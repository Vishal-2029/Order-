package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	Items      string  `json:"items"`
	TotalPrice float64 `json:"totalPrice"`
	Status     string  `json:"status"`
}
