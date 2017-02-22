package controllers

import (
	"com_web/models"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type NewsController struct {
	beego.Controller
}

func (c *NewsController) Get() {
	db, _ := gorm.Open("mysql", "root:special005@/com_admin?charset=utf8&parseTime=True&loc=Local")
	var (
		blogs models.Blog
	)
	blog := db.Where(&models.Blog{Author: "root"}).First(&blogs).Value
	c.Data["time"] = blog.(*models.Blog).CreatedAt.Format("2006-01-02")
	c.Data["title"] = blog.(*models.Blog).Title
	c.Data["content"] = blog.(*models.Blog).Content
	c.Data["author"] = blog.(*models.Blog).Author
	c.TplName = "blog.tpl"
}
