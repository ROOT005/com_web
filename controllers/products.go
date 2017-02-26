package controllers

import (
	//"com_web/db"
	"com_web/models"
	"github.com/astaxie/beego"
	"strconv"
)

type ProductsController struct {
	beego.Controller
}

func (c *ProductsController) Get() {

	//products, err := models.GetAllProducts()
	/****************传入数据**************/
	pre_page := 10
	pa := 1
	pa, _ = strconv.Atoi(c.Input().Get("p"))
	products := models.GetAllProducts(pa, pre_page)
	c.Data["paginator"] = products
	/*if err {
		c.Data["products"] = products
	} else {
		c.Data["products"] = nil
	}*/
	c.Data["head_title"] = "产品中心"
	c.Data["style_name"] = "product"
	c.TplName = "products.tpl"
}
