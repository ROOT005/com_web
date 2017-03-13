package controllers

import (
	"com_web/models"
	"encoding/csv"
	"github.com/astaxie/beego"
	"os"
	"strconv"
	"time"
)

type NewUsersController struct {
	beego.Controller
}
type DowloadController struct {
	beego.Controller
}

func (c *NewUsersController) Get() {
	id := c.Input().Get("id")
	newusers := models.GetNewUsers(id)
	c.Ctx.Output.JSON(newusers, true, true)
}

func (c *DowloadController) Get() {
	filename := "data/" + time.Now().Format("2006-01-02") + ".csv"
	f, err := os.Create(filename)
	if err != nil {
		c.Ctx.WriteString("下载失败!")
	}
	//取数据
	id := "1"
	users := models.GetNewUsers(id)
	//导出数据
	defer f.Close()
	defer os.RemoveAll(filename)
	f.WriteString("\xEF\xBB\xBF") //声明编码格式 utf-8
	w := csv.NewWriter(f)
	w.Write([]string{"ID", "姓名", "需求", "电话", "时间"})
	for _, user := range users {
		w.Write([]string{strconv.Itoa(int(user.ID)), user.Name, user.Count, user.Phone, user.CreatedAt.Format("2006-01-02 15:04:05")})
	}
	w.Flush()
	c.Ctx.Output.Download(filename)
}
