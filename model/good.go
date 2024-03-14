package model

import (
	"gorm.io/gorm"
)

// 商品表
type CGoods struct {
	gorm.Model
	GoodsName   string `gorm:"index"`             //商品名
	Image       string `gorm:"type:text(0)"`      //头图
	Description string `gorm:"type:varchar(250)"` //描述
	Brand       int64  `gorm:"type:varchar(100)"` //品牌
	Status      int64  `gorm:"type:tinyint(1)"`   //状态
	Supplier    string `gorm:"type:varchar(255)"` //供应商
}

// 商品属性表
type Property struct {
	gorm.Model
	GoodsId int64  `gorm:"index,type:int(10)"` //商品id
	Name    string `gorm:"type:varchar(100)"`  //商品属性名
}

// 商品属性值
type PropertyValue struct {
	gorm.Model
	PropertyId int64  `gorm:"index,type:int(10)"` //属性id
	Value      string `gorm:"type:varchar(100)"`  //属性值
}

// 商品sku
type Sku struct {
	gorm.Model
	GoodsId     int64  `gorm:"index,type:int(10)"` //商品id
	ProValIdFir int64  `gorm:"index,type:int(10)"` //第一个商品属性值
	ProValIdSec int64  `gorm:"index,type:int(10)"` //第二个商品属性值
	ProValIdThi int64  `gorm:"index,type:int(10)"` //第三个商品属性值
	Price       string `gorm:"type:decimal(10,2)"` //单价
	Stock       int64  `gorm:"index,type:int(6)"`  //库存
}

func NewCGoods() *CGoods {
	return new(CGoods)
}

func MewProperty() *Property {
	return new(Property)
}

func MewPropertyValue() *PropertyValue {
	return new(PropertyValue)
}

func MewSku() *Sku {
	return new(Sku)
}

func (g *CGoods) Create(goods *CGoods) (info *CGoods, err error) {
	err = db.Model(g).Create(&goods).Error
	info = goods
	return info, err
}

func (g *Property) Create(goods *Property) (info *Property, err error) {
	err = db.Model(g).Create(&goods).Error
	info = goods
	return info, err
}

func (g *PropertyValue) Create(goods *PropertyValue) (info *PropertyValue, err error) {
	err = db.Model(g).Create(&goods).Error
	info = goods
	return info, err
}

func (g *Sku) Create(goods *Sku) (info *Sku, err error) {
	err = db.Model(g).Create(&goods).Error
	info = goods
	return info, err
}
