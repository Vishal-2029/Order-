package models

import "gorm.io/gorm"

type Items struct {
	gorm.Model
	ItemName string  `json:"item_name"`
	Price    float64 `json:"price"`
	Stock    int     `json:"stock"`
}
