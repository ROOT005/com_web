package controllers

import (
	"com_web/db"
	"com_web/models"
	"github.com/astaxie/beego"
)

type ProductsController struct {
	beego.Controller
}

func (c *ProductsController) Get() {
	var (
		products models.Product
	)
	product := db.DB.First(&products).Limit(1).Value
	c.Data["product"] = product.(*models.Product)
	c.Data["head_title"] = "产品中心"
	c.Data["style_name"] = "product"
	c.TplName = "products.tpl"
}
