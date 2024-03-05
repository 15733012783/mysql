package model

import (
	"fmt"
	"github.com/15733012783/mysql/nacos"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB
var err error

func InItMysql() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", nacos.NaCosT.Username,
		nacos.NaCosT.Password, nacos.NaCosT.Host, nacos.NaCosT.Port, nacos.NaCosT.Mysqlbase)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	err := db.AutoMigrate(new(Goods))
	if err != nil {
		log.Println(err, "**********************AutoMigrate")
		return
	}
}
