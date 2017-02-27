package controllers

import (
	"com_web/models"
	"github.com/astaxie/beego"
	"strconv"
)

type ProjectController struct {
	beego.Controller
}

func (c *ProjectController) Get() {
	/*传入数据*/
	pre_page := 10
	pa := 1
	var (
		category     int
		projectrange int
	)
	//获取参数
	pa, _ = strconv.Atoi(c.Input().Get("p"))
	category, _ = strconv.Atoi(c.Input().Get("category"))
	projectrange, _ = strconv.Atoi(c.Input().Get("projectrange"))
	//获取数据库内容
	projects := models.GetAllProject(pa, pre_page, category, projectrange)
	categories := models.GetAllCategory()
	projectranges := models.GetAllProductRange()
	c.Data["projectranges"] = projectranges
	c.Data["paginator"] = projects
	c.Data["categories"] = categories
	c.Data["head_title"] = "项目中心"
	c.Data["style_name"] = "project"
	c.TplName = "projects.tpl"
}

//获取指定项目信息
func (c *ProjectController) Project_Info() {
	id := c.Input().Get("id")
	num, _ := strconv.Atoi(id)
	project_info := models.GetProject(num)[0]
	c.Ctx.Output.JSON(project_info, true, true)
}
