package controllers

import (
	//"com_web/db"
	"com_web/models"
	"github.com/astaxie/beego"
)

type ProductsController struct {
	beego.Controller
}

func (c *ProductsController) Get() {

	//products, err := models.GetAllProducts()
	products := models.GetAllProducts()
	c.Data["products"] = products
	/*if err {
		c.Data["products"] = products
	} else {
		c.Data["products"] = nil
	}*/
	c.Data["head_title"] = "产品中心"
	c.Data["style_name"] = "product"
	c.TplName = "products.tpl"
}
