package model

import (
	"fmt"
	"github.com/15733012783/mysql/nacos"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// var db *gorm.DB
// var err error
func inItMysql(c func(db *gorm.DB) (interface{}, error)) (interface{}, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		nacos.GoodsT.Mysql.Username,
		nacos.GoodsT.Mysql.Password,
		nacos.GoodsT.Mysql.Host,
		nacos.GoodsT.Mysql.Port,
		nacos.GoodsT.Mysql.Mysqlbase)
	open, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}
	db, err := open.DB()
	if err != nil {
		return nil, err
	}
	s, err := c(open)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	return s, err
}

func InitTable() {
	inItMysql(func(db *gorm.DB) (interface{}, error) {
		err := db.AutoMigrate(new(Goods), new(Property), new(PropertyValue), new(Sku))
		if err != nil {
			return nil, err
		}
		return nil, nil
	})
}
