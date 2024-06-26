package model

import (
	"fmt"
	"github.com/15733012783/mysql/nacos"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func InItMysql() {
	var db *gorm.DB
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		nacos.GoodsT.Mysql.Username,
		nacos.GoodsT.Mysql.Password,
		nacos.GoodsT.Mysql.Host,
		nacos.GoodsT.Mysql.Port,
		nacos.GoodsT.Mysql.Mysqlbase)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	err = db.AutoMigrate(new(CGoods))
	if err != nil {
		log.Println(err)
		return
	}
}
