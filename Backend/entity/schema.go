package entity

import (
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Name       string `valid:"required~name canonot be blank"`
	Email      string
	CustomerID string `valid:"matches(^[LMD]\\$d[7])~not valid`
}
