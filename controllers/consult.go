package controllers

import (
	"com_web/models"
	//"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/utils/captcha"
	"github.com/astaxie/beego/validation"
	"strconv"
)

type SubmitController struct {
	beego.Controller
}

func (c *SubmitController) Get() {
	c.Redirect("/home", 301)
}
func (c *SubmitController) Post() {
	name := c.Input().Get("name")
	demend := c.Input().Get("demend")
	count := c.Input().Get("count")
	phone := c.Input().Get("phone")
	others := c.Input().Get("others")
	err := models.AddUser(name, demend, count, phone, others)
	if err != nil {
		beego.Error(err)
	}
	c.Redirect("/", 302)
	return
}
func (c *SubmitController) Phone() {
	phone := c.Input().Get("phone")
	valid := validation.Validation{}
	if valid.Phone(phone, "Mobile").Ok && valid.Required(phone, "Mobile").Ok {
		c.Ctx.WriteString("true")
	} else {
		c.Ctx.WriteString("false")
	}
}

//数据验证
type VerifyController struct {
	beego.Controller
}

var cpt *captcha.Captcha

func init() {
	//验证码
	var store = cache.NewMemoryCache()
	cpt = captcha.NewWithFilter("/captcha/", store)
	cpt.ChallengeNums = 4
	cpt.StdWidth = 200
	cpt.StdHeight = 34
}
func (c *VerifyController) Get() {
	result := cpt.Verify(c.Input().Get("captcha_id"), c.Input().Get("id"))
	c.Ctx.WriteString(strconv.FormatBool(result))
}

//信息页
type ConsultController struct {
	beego.Controller
}

func (c *ConsultController) Get() {
	c.TplName = "c_info.tpl"
}
