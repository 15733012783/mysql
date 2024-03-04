package model

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB
var err error

func InItMysql() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", NaCosT.Username,
		NaCosT.Password, NaCosT.Host, NaCosT.Port, NaCosT.Mysqlbase)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	err := db.AutoMigrate(new(Goods))
	if err != nil {
		log.Println(err, "**********************AutoMigrate")
		return
	}
}
