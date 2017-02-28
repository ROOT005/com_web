package controllers

import (
	"com_web/db"
	"com_web/models"
	"github.com/astaxie/beego"
	//"github.com/jinzhu/gorm"
)

type NewsController struct {
	beego.Controller
}

func (c *NewsController) Get() {
	var (
		blogs models.Blog
	)
	category := models.GetAllBlogCategory()
	blog := db.DB.Where(&models.Blog{Author: "root"}).First(&blogs).Value
	c.Data["time"] = blog.(*models.Blog).CreatedAt.Format("2006-01-02")
	c.Data["title"] = blog.(*models.Blog).Title
	c.Data["abstract"] = blog.(*models.Blog).Abstract
	c.Data["author"] = blog.(*models.Blog).Author
	c.Data["head_title"] = "新闻动态"
	c.Data["style_name"] = "blog"
	c.Data["blogcategory"] = category
	c.TplName = "blog.tpl"
}
