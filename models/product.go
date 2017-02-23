package models

import (
	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	Name          string
	Count         string
	Rate          string
	RepayTime     string
	RepayWay      string
	SourceCompany string
	Industry      string
	LoanTime      string
	Description   string `sql:"size:5000"`
}
