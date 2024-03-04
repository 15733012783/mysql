package model

import (
	"fmt"
	"gorm.io/gorm"
)

type Goods struct {
	gorm.Model
	GoodsName  string  `gorm:"index"`
	GoodsPrice float64 `gorm:"decimal(10,2)"`
	GoodsNum   int     `gorm:"tinyint"`
	GoodsPhoto string  `gorm:"varchar(100)"`
}

func NewGoods() *Goods {
	return new(Goods)
}

func (g *Goods) Create(goods *Goods) (info *Goods, err error) {
	err = db.Model(g).Create(&goods).Error
	info = goods
	return info, err
}

func (g *Goods) Delete(id int) error {
	err := db.Model(g).Where("id = ?", id).Delete(g).Error
	return err
}

func (g *Goods) Upload(goods *Goods) (info *Goods, err error) {
	err = db.Model(g).Where("id = ?", goods.ID).Updates(goods).Error
	return goods, err
}

func (g *Goods) Get(goodsName string) (info *Goods, err error) {
	info = new(Goods)
	err = db.Model(g).Where("goods_name = ?", goodsName).First(&info).Error
	return info, err
}

func (g *Goods) UploadFile(id int, FileName string) (info *Goods, err error) {
	info = new(Goods)
	err = db.Model(g).Where("id = ?", id).First(info).Error
	fmt.Println("**********************************************info")
	fmt.Println(info)
	info.GoodsPhoto = FileName
	err = db.Model(g).Save(info).Error
	fmt.Println("**********************************************info")
	fmt.Println(info)
	return info, err
}
