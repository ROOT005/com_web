package controllers

import (
	"com_web/models"
	"fmt"
	"github.com/astaxie/beego"
	"reflect"
	"strconv"
)

type NewsController struct {
	beego.Controller
}

func (c *NewsController) Get() {
	pre_page := 10
	pa := 1
	var (
		category int
	)
	pa, _ = strconv.Atoi(c.Input().Get("p"))
	category, _ = strconv.Atoi(c.Input().Get("blogcategory"))
	//获取数据库内容
	blogs := models.GetAllBlog(pa, pre_page, category)
	categories := models.GetAllBlogCategory()
	c.Data["blogcategory"] = categories
	c.Data["paginator"] = blogs
	fmt.Println(reflect.TypeOf(blogs["blogs"]))
	for i, k := range blogs["blogs"].([]models.Blog) {
		fmt.Println(i, k)
	}
	c.Data["head_title"] = "新闻动态"
	c.Data["style_name"] = "blog"
	c.TplName = "blog.tpl"
}
