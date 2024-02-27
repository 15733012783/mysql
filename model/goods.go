package model

import (
	"github.com/15733012783/proto"
	"gorm.io/gorm"
)

type Goods struct {
	gorm.Model
	GoodsName  string
	GoodsPrice float64
	GoodsNum   int
}

func NewGoods() *Goods {
	return new(Goods)
}

func (g *Goods) Create(goods Goods) (info *proto.GoodsInfo, err error) {
	err = db.Model(g).Create(&goods).Error
	return info, err
}

func (g *Goods) Delete(id int) error {
	err := db.Model(g).Where("id = ?", id).Delete(g).Error
	return err
}

func (g *Goods) Upload(goods Goods) (info Goods, err error) {
	err = db.Model(g).Where("id = ?", goods.ID).Updates(goods).Error
	return goods, err
}

func (g *Goods) Get(id int) (info *proto.GoodsInfo, err error) {
	var goods Goods
	err = db.Model(g).Where("id = ?", id).First(&goods).Error
	return info, err
}
