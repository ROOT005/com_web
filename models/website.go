package models

import (
	"github.com/jinzhu/gorm"
	"github.com/qor/media_library"
	qor_seo "github.com/qor/seo"
)

/*文章*/
type Blog struct {
	gorm.Model
	KeyWords string
	Author   string
	Title    string
	Content  string `sql:"size:5000"`
	//Seo     qor_seo.Setting
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
