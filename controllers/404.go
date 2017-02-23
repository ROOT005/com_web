package controllers

import (
	"github.com/astaxie/beego"
)

type ErrorController struct {
	beego.Controller
}

func (e *ErrorController) Error404() {
	e.TplName = "404.html"
}
