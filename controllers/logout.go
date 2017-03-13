package controllers

import (
	"github.com/astaxie/beego"
)

type LogoutController struct {
	beego.Controller
}

func (c *LogoutController) Get() {
	c.Ctx.SetCookie("id", "", -1, "/")
	c.Redirect("/login", 301)
}
