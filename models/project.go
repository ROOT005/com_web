package models

import (
	"com_web/db"
	"github.com/jinzhu/gorm"
	"math"
)

type Category struct {
	gorm.Model
	Name string
}
type ProjectRange struct {
	gorm.Model
	Name string
}
type Project struct {
	gorm.Model
	Name           string
	Count          string
	Category       Category
	CategoryID     uint
	ProjectRange   ProjectRange
	ProjectRangeID uint
	Industry       string
	FinanceCompany string
	Location       string
	FinanceWay     string
	Description    string `sql:"size:5000"`
}

//获取分类
func GetAllCategory() []*Category {
	var category []*Category
	db.DB.Find(&category)
	return category
}

//获取需求范围

func GetAllProductRange() []*ProjectRange {
	var projectrange []*ProjectRange
	db.DB.Find(&projectrange)
	return projectrange
}

//获取指定项目信息

func GetProject(id int) []*Project {
	var project []*Project
	db.DB.First(&project, id)
	return project
}

//获取项目
func GetAllProject(page, prepage int, category, projectrange int) map[string]interface{} {
	var project []*Project
	if category == 0 && projectrange == 0 {
		db.DB.Find(&project)
	}
	if category != 0 && projectrange == 0 {
		db.DB.Where("category_id = ?", category).Find(&project)
	}
	if category == 0 && projectrange != 0 {
		db.DB.Where("project_range_id = ? ", projectrange).Find(&project)
	}
	if category != 0 && projectrange != 0 {
		db.DB.Where(map[string]interface{}{
			"categort_id":      category,
			"project_range_id": projectrange,
		}).Find(&project)
	}
	/********************分页*******************************/
	nums := len(project)
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
	var projects []Project
	startP := (page - 1) * prepage
	if category == 0 && projectrange == 0 {
		db.DB.Limit(prepage).Offset(startP).Find(&projects)
	}
	if category != 0 && projectrange == 0 {
		db.DB.Where("category_id = ?", category).Limit(prepage).Offset(startP).Find(&projects)
	}
	if category == 0 && projectrange != 0 {
		db.DB.Where("project_range_id = ? ", projectrange).Limit(prepage).Offset(startP).Find(&projects)
	}
	if category != 0 && projectrange != 0 {
		db.DB.Where(map[string]interface{}{
			"categort_id":      category,
			"project_range_id": projectrange,
		}).Limit(prepage).Offset(startP).Find(&projects)
	}
	paginatorMap := make(map[string]interface{})
	paginatorMap["projects"] = projects
	paginatorMap["pages"] = pages
	paginatorMap["totalpages"] = totalpages
	paginatorMap["firstpage"] = firstpage
	paginatorMap["lastpage"] = lastpage
	paginatorMap["currpage"] = page
	paginatorMap["totals"] = nums
	return paginatorMap
}
