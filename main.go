package main

import (
	_ "com_web/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.SetStaticPath("/system", "public/system")
	beego.SetStaticPath("/admin/asset", "public/admin/asset")
	beego.Run()
}
