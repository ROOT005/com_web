package models

import (
	"com_web/db"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name   string
	Demend string
	Count  string
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

	var user = User{Name: name, Demend: demendc, Count: count, Phone: phone, Others: others}
	return db.DB.Create(&user).Error
}

func (user AdminUser) DisplayName() string {
	return user.Email
}
