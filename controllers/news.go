package controllers

import (
	"com_web/models"
	"github.com/astaxie/beego"
	"strconv"
)

type NewsController struct {
	beego.Controller
}

func (c *NewsController) Get() {
	pre_page := 8
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
	c.Data["head_title"] = "新闻动态"
	c.Data["style_name"] = "blog"
	c.TplName = "blog.tpl"
}

//获取详情
func (c *NewsController) News_Info() {
	id := c.Input().Get("id")
	num, _ := strconv.Atoi(id)
	news_info := models.GetBlog(num)
	c.Data["new_info"] = news_info
	c.TplName = "news_info.html"
}
