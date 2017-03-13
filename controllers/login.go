package controllers

import (
	"com_web/db"
	"com_web/models"
	"github.com/astaxie/beego"
	"golang.org/x/crypto/bcrypt"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	c.TplName = "login.html"
}
func (c *LoginController) Post() {
	email := c.Input().Get("email")
	value := "BgQDwQ3THJn9F7NPLBi6hTI3Fwz55h47jQUVCOL6iq"
	password := c.Input().Get("password")
	var adminuser models.AdminUser
	if !db.DB.First(&adminuser, "email = ?", c.Input().Get("email")).RecordNotFound() {
		var passwd string
		passwd = adminuser.Password
		err := bcrypt.CompareHashAndPassword([]byte(passwd), []byte(password))
		if err == nil {
			c.Ctx.SetCookie("id", email, "/")
			c.Ctx.SetCookie("see", value, "/")
			c.Redirect("/admin", 301)
			return
		} else {
			c.Redirect("/", 301)
		}

	} else {
		c.Redirect("/", 301)
	}
}
