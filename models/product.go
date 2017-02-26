package models

import (
	"com_web/db"
	//"fmt"
	"github.com/jinzhu/gorm"
	"math"
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

func GetAllProducts(page, prepage int) map[string]interface{} {
	var product []Product
	db.DB.Find(&product)

	/*****************分页********************/
	nums := len(product)
	var firstpage int //前一页地址
	var lastpage int  //后一页地址

	//根据nums总数，和prepage每页数量 生成分页总数
	totalpages := int(math.Ceil(float64(nums) / float64(prepage))) //page总数
	if page > totalpages {
		page = totalpages
	}
	if page <= 0 {
		page = 1
	}
	var pages []int //所有页
	switch {
	case page >= totalpages-5 && totalpages > 5: //最后5页
		start := totalpages - 5 + 1
		firstpage = page - 1
		lastpage = int(math.Min(float64(totalpages), float64(page+1)))
		pages = make([]int, 5)
		for i, _ := range pages {
			pages[i] = start + i
		}
	case page >= 3 && totalpages > 5:
		start := page - 3 + 1
		pages = make([]int, 5)
		firstpage = page - 3
		for i, _ := range pages {
			pages[i] = start + i
		}
		firstpage = page - 1
		lastpage = page + 1
	default:
		pages = make([]int, int(math.Min(5, float64(totalpages))))
		for i, _ := range pages {
			pages[i] = i + 1
		}
		firstpage = int(math.Max(float64(1), float64(page-1)))
		lastpage = page + 1
	}
	//获取数据
	var products []Product
	startP := (page - 1) * prepage

	db.DB.Find(&products).Limit(prepage).Offset(startP)

	paginatorMap := make(map[string]interface{})
	paginatorMap["products"] = products
	paginatorMap["pages"] = pages
	paginatorMap["totalpages"] = totalpages
	paginatorMap["firstpage"] = firstpage
	paginatorMap["lastpage"] = lastpage
	paginatorMap["currpage"] = page
	paginatorMap["totals"] = nums
	return paginatorMap
}
