package models

import (
	"com_web/db"
	"github.com/jinzhu/gorm"
	"github.com/qor/media_library"
	qor_seo "github.com/qor/seo"
	"math"
)

/*文章*/
type Blog struct {
	gorm.Model
	KeyWords       string
	BlogCategory   BlogCategory
	BlogCategoryID uint
	Author         string
	Title          string
	Mainimage      media_library.MediaBox
	Abstract       string `sql:"size:5000"`
	Content        string `sql:"size:10000"`
	//Seo     qor_seo.Setting
}

/*博客分类*/
type BlogCategory struct {
	gorm.Model
	Name string
}

//获取文章
func GetBlog(id int) []*Blog {
	var blog []*Blog
	db.DB.First(&blog, id)
	return blog
}

//获取文章分类
func GetAllBlogCategory() []*BlogCategory {
	var blogcategory []*BlogCategory
	db.DB.Find(&blogcategory)
	return blogcategory
}

//获取全部文章
func GetAllBlog(page, prepage, blogcategory int) map[string]interface{} {
	var blog []*Blog

	if blogcategory == 0 {
		db.DB.Find(&blog)
	}

	if blogcategory != 0 {
		db.DB.Where("blog_category_id", blogcategory).Find(&blog)
	}

	/*****************分页********************/
	nums := len(blog)
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
	var blogs []Blog
	startP := (page - 1) * prepage

	if blogcategory != 0 {
		db.DB.Where("blog_category_id = ?", blogcategory).Limit(prepage).Offset(startP).Find(&blogs)
	}
	if blogcategory == 0 {
		db.DB.Limit(prepage).Offset(startP).Find(&blogs)
	}

	paginatorMap := make(map[string]interface{})
	paginatorMap["blogs"] = blogs
	paginatorMap["pages"] = pages
	paginatorMap["totalpages"] = totalpages
	paginatorMap["firstpage"] = firstpage
	paginatorMap["lastpage"] = lastpage
	paginatorMap["currpage"] = page
	paginatorMap["totals"] = nums
	return paginatorMap
}

/*主页内容*/
type Index struct {
	gorm.Model
	Title      string
	LSignWords string
	RSignWords string
	Logo       media_library.MediaBox
	Seo        qor_seo.Setting
}
type Pictures struct {
	gorm.Model
	Title string
}
type Seo struct {
	gorm.Model
	SiteName    string
	Description string
	KeyWords    string
}
