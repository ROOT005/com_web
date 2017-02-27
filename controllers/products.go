package controllers

import (
	"com_web/models"
	//"fmt"
	"github.com/astaxie/beego"
	"strconv"
)

type ProductsController struct {
	beego.Controller
}

func (c *ProductsController) Get() {

	/****************传入数据**************/
	pre_page := 10
	pa := 1
	var class int
	var rangec int

	pa, _ = strconv.Atoi(c.Input().Get("p"))
	class, _ = strconv.Atoi(c.Input().Get("id"))
	rangec, _ = strconv.Atoi(c.Input().Get("rangec"))
	products := models.GetAllProducts(pa, pre_page, class, rangec)
	classfication := models.GetAllClassfication()
	rangecount := models.GetAllRange()

	c.Data["rangecount"] = rangecount
	c.Data["paginator"] = products
	c.Data["classfication"] = classfication
	c.Data["head_title"] = "产品中心"
	c.Data["style_name"] = "product"
	c.TplName = "products.tpl"
}

//获取指定产品信息
func (c *ProductsController) Product_Info() {
	id := c.Input().Get("id")
	num, _ := strconv.Atoi(id)
	product_info := models.GetProduct(num)[0]
	c.Ctx.Output.JSON(product_info, true, true)
}
