package config

import (
	"fmt"
	"github.com/FadhlanHawali/Digitalent-Kominfo_Implementation-MVC-Golang/app/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DBInit() *gorm.DB{
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("root:pintar123@/digitalent_bank?charset=utf8&parseTime=True&loc=Local")), &gorm.Config{})
	if err != nil {
		panic("failede to connect to database" + err.Error())
	}
	db.AutoMigrate(new(model.Account),new(model.Transaction))
	return db
}
