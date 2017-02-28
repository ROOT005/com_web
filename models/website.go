package models

import (
	"com_web/db"
	"github.com/jinzhu/gorm"
	"github.com/qor/media_library"
	qor_seo "github.com/qor/seo"
)

/*文章*/
type Blog struct {
	gorm.Model
	KeyWords       string
	BlogCategory   BlogCategory
	BlogCategoryID uint
	Author         string
	Title          string
	Abstract       string `sql:"size:1000`
	Content        string `sql:"size:5000"`
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
func GetAllBlog() {

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
