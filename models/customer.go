package models

import (


	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	CustomerName string    `json:"customerName"`
	Passcode     string    `json:"passcode"`
}
