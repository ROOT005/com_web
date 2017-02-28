package controllers

import (
	"github.com/astaxie/beego"
)

type ConsultController struct {
	beego.Controller
}

func (c *ConsultController) Get() {
	c.TplName = "c_info.tpl"
}
