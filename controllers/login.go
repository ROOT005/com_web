package controllers

import (
	"com_web/db"
	"com_web/models"
	"fmt"
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
	password := c.Input().Get("password")
	//cryptpasswd, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	var adminuser models.AdminUser
	if !db.DB.First(&adminuser, "email = ?", email).RecordNotFound() {
		var passwd string
		passwd = adminuser.Password
		err := bcrypt.CompareHashAndPassword([]byte(passwd), []byte(password))
		if err == nil {
			fmt.Print("hello\n")
			c.Ctx.SetCookie("id", email, "/")
			c.Redirect("/admin", 301)
			return
		} else {
			c.Redirect("/", 301)
		}

	} else {
		c.Redirect("/", 301)
	}
}
