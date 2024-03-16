package model

import (
	"fmt"
	"gorm.io/gorm"
)

type CGoods struct {
	gorm.Model
	GoodsName  string  `gorm:"index"`
	GoodsPrice float64 `gorm:"decimal(10,2)"`
	GoodsNum   int     `gorm:"tinyint"`
	GoodsPhoto string  `gorm:"varchar(100)"`
}

func NewGoods() *CGoods {
	return new(CGoods)
}

func (g *CGoods) Create(goods *CGoods) (info *CGoods, err error) {
	err = db.Model(g).Create(&goods).Error
	info = goods
	return info, err
}

func (g *CGoods) Delete(id int) error {
	err := db.Model(g).Where("id = ?", id).Delete(g).Error
	return err
}

func (g *CGoods) Upload(goods *CGoods) (info *CGoods, err error) {
	err = db.Model(g).Where("id = ?", goods.ID).Updates(goods).Error
	return goods, err
}

func (g *CGoods) Get(goodsName string) (info *CGoods, err error) {
	info = new(CGoods)
	err = db.Model(g).Where("goods_name = ?", goodsName).First(&info).Error
	return info, err
}

func (g *CGoods) UploadFile(id int, FileName string) (info *CGoods, err error) {
	info = new(CGoods)
	err = db.Model(g).Where("id = ?", id).First(info).Error
	fmt.Println("**********************************************info")
	fmt.Println(info)
	info.GoodsPhoto = FileName
	err = db.Model(g).Save(info).Error
	fmt.Println("**********************************************info")
	fmt.Println(info)
	return info, err
}
