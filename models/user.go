package models

import (
	"com_web/db"
	"github.com/jinzhu/gorm"
	"strconv"
)

type User struct {
	gorm.Model
	Name   string
	Demend string
	Count  string
	State  string
	Phone  string
	Others string
}
type AdminUser struct {
	gorm.Model
	Email    string
	Password string
	Name     string
	Role     string
}

func AddUser(name, demend, count, phone, others string) error {
	var demendc string
	switch demend {
	case "1":
		demendc = "找资金"
	case "2":
		demendc = "资方入驻"
	}

	var user = User{Name: name, Demend: demendc, Count: count, State: "未查看", Phone: phone, Others: others}
	return db.DB.Create(&user).Error
}
func (user AdminUser) DisplayName() string {
	return user.Email
}

func GetNewUsers(idn string) []*User {
	id, _ := strconv.Atoi(idn)
	var newnusers []*User
	db.DB.Where("id > ? AND demend = ? AND state = ?", id, "找资金", "未查看").Order("created_at desc").Find(&newnusers)
	return newnusers
}
