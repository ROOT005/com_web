package models

import (
	"com_web/db"
	//"fmt"
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

/*func GetAllProducts() ([]*Product, bool) {
	var product Product
	products := make([]*Product, 0)
	allProducts, err := db.DB.Find(&product).Value.(*Product)
	products = append(products, allProducts)
	fmt.Println(db.DB.Find(&product).Value/*)
	return products, err
}*/
func GetAllProducts() []Product {
	var product []Product
	db.DB.Find(&product)
	return product
}
