package model

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var db *gorm.DB
var err error

func InItMysql() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", NaCosT.Username,
		NaCosT.Password, NaCosT.Host, NaCosT.Port, NaCosT.Mysqlbase)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	MysqlDate()
}

func MysqlDate() {
	err := db.AutoMigrate(new(User))
	if err != nil {
		log.Println(err)
		return
	}
}
