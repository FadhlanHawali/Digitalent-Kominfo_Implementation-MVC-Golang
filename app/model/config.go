package model

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
func init () {
	var err error
	DB, err = gorm.Open(mysql.Open(fmt.Sprintf("root:pintar123@/digitalent_bank?charset=utf8&parseTime=True&loc=Local")), &gorm.Config{})
	if err != nil {
		panic("failede to connect to database" + err.Error())
	}
	DB.AutoMigrate(new(Account),new(Transaction))

}
