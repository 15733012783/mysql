package model

import (
	"gorm.io/gorm"
)

// 商品表
type Goods struct {
	gorm.Model
	GoodsName   string `gorm:"index"`             //商品名
	Image       string `gorm:"type:text(0)"`      //头图
	Description string `gorm:"type:varchar(250)"` //描述
	Brand       string `gorm:"type:varchar(100)"` //品牌
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
	PropertyValueID int64  `gorm:"index,type:int(10)"`
	Price           string `gorm:"type:decimal(10,2)"` //单价
	Stock           int64  `gorm:"index,type:int(6)"`  //库存
	Discount        int64  `gorm:"index,type:int(6)"`  //折扣
	State           int64  `gorm:"index,type:int(6)"`
	Form            string `gorm:"type:varchar(100)"`
}

func NewCGoods() *Goods {
	return new(Goods)
}

func NewProperty() *Property {
	return new(Property)
}

func NewPropertyValue() *PropertyValue {
	return new(PropertyValue)
}

func NewSku() *Sku {
	return new(Sku)
}

func (g *Goods) Create(goods *Goods) (info *Goods, err error) {
	itMysql, err := inItMysql(func(db *gorm.DB) (interface{}, error) {
		tx := db.Begin()
		err = tx.Model(&Goods{}).Create(&goods).Error
		if err != nil {
			tx.Rollback()
		}
		info = goods
		tx.Commit()
		return info, err
	})
	inM := itMysql.(*Goods)
	return inM, err
}

func (g *Goods) Delete(id int) error {
	_, err2 := inItMysql(func(db *gorm.DB) (interface{}, error) {
		tx := db.Begin()
		err = tx.Model(&Goods{}).Delete("id = ?", id).Error
		if err != nil {
			tx.Rollback()
		}
		tx.Commit()
		return nil, err
	})
	return err2
}

func (g *Property) Create(goods *Property) (info *Property, err error) {
	itMysql, err := inItMysql(func(db *gorm.DB) (interface{}, error) {
		tx := db.Begin()
		err = tx.Model(&Property{}).Create(&goods).Error
		if err != nil {
			tx.Rollback()
		}
		info = goods
		tx.Commit()
		return info, err
	})
	inM := itMysql.(*Property)
	return inM, err
}

func (g *Goods) Where(page, num int64) (info []*Goods, err error) {
	itMysql, err := inItMysql(func(db *gorm.DB) (interface{}, error) {
		var goods []*Goods
		offset := (page - 1) * num
		err = db.Model(&Goods{}).Offset(int(offset)).Limit(int(num)).Find(&goods).Error
		if err != nil {
			return nil, err
		}
		return goods, err
	})
	inM := itMysql.([]*Goods)
	return inM, err
}

func (g *Property) WhereID(name string) (info *Property, err error) {
	itMysql, err := inItMysql(func(db *gorm.DB) (interface{}, error) {
		err = db.Model(&Property{}).Where("name = ?", name).First(&info).Error
		return info, err
	})
	inM := itMysql.(*Property)
	return inM, err
}

func (g *Property) Where(ID int) (info []*Property, err error) {
	itMysql, err := inItMysql(func(db *gorm.DB) (interface{}, error) {
		var property []*Property
		err = db.Model(&Property{}).Where("goods_id = ?", ID).Find(&property).Error
		return property, err
	})
	inM := itMysql.([]*Property)
	return inM, err
}

func (g *Property) Delete(id int) error {
	_, err2 := inItMysql(func(db *gorm.DB) (interface{}, error) {
		tx := db.Begin()
		err = tx.Model(&Property{}).Delete("id = ?", id).Error
		if err != nil {
			tx.Rollback()
		}
		tx.Commit()
		return nil, err
	})
	return err2
}

func (g *PropertyValue) Create(goods *PropertyValue) (info *PropertyValue, err error) {
	itMysql, err := inItMysql(func(db *gorm.DB) (interface{}, error) {
		tx := db.Begin()
		err = tx.Model(&PropertyValue{}).Create(&goods).Error
		if err != nil {
			tx.Rollback()
		}
		info = goods
		tx.Commit()
		return info, err
	})
	inM := itMysql.(*PropertyValue)
	return inM, err
}

func (g *PropertyValue) WhereID(id int) (info *PropertyValue, err error) {
	itMysql, err := inItMysql(func(db *gorm.DB) (interface{}, error) {
		err = db.Model(&PropertyValue{}).Where("property_id = ?", id).First(&info).Error
		return info, err
	})
	inM := itMysql.(*PropertyValue)
	return inM, err
}

func (g *PropertyValue) Where(ID int) (info []*PropertyValue, err error) {
	itMysql, err := inItMysql(func(db *gorm.DB) (interface{}, error) {
		var propertyValue []*PropertyValue
		err = db.Model(&PropertyValue{}).Where("property_id = ?", ID).Find(&propertyValue).Error
		return propertyValue, err
	})
	inM := itMysql.([]*PropertyValue)
	return inM, err
}

func (g *PropertyValue) Delete(id int) error {
	_, err2 := inItMysql(func(db *gorm.DB) (interface{}, error) {
		tx := db.Begin()
		err = tx.Model(&PropertyValue{}).Delete("id = ?", id).Error
		if err != nil {
			tx.Rollback()
		}
		tx.Commit()
		return nil, err
	})
	return err2
}

func (g *Sku) Create(goods *Sku) (info *Sku, err error) {
	itMysql, err := inItMysql(func(db *gorm.DB) (interface{}, error) {
		tx := db.Begin()
		err = db.Model(&Sku{}).Create(&goods).Error
		if err != nil {
			tx.Rollback()
		}
		info = goods
		tx.Commit()
		return info, err
	})
	inM := itMysql.(*Sku)
	return inM, err
}

func (g *Sku) Where(ID int) (info []*Sku, err error) {
	itMysql, err := inItMysql(func(db *gorm.DB) (interface{}, error) {
		var sku []*Sku
		err = db.Model(&Sku{}).Where("property_value_id = ?", ID).Find(&sku).Error
		return sku, err
	})
	inM := itMysql.([]*Sku)
	return inM, err
}

func (g *Sku) Delete(id int) error {
	_, err2 := inItMysql(func(db *gorm.DB) (interface{}, error) {
		tx := db.Begin()
		err = tx.Model(&Sku{}).Delete("id = ?", id).Error
		if err != nil {
			tx.Rollback()
		}
		tx.Commit()
		return nil, err
	})
	return err2
}
