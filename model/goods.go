package model

import "gorm.io/gorm"

type Goods struct {
	gorm.Model
	GoodsName  string
	GoodsPrice float64
	GoodsNum   int
}

func NewGoods() *Goods {
	return new(Goods)
}

func (g *Goods) Create(goods Goods) (info *Goods, err error) {
	err = db.Model(g).Create(&goods).Error
	info = &goods
	return info, err
}

func (g *Goods) Delete(id int) error {
	err := db.Model(g).Where("id = ?", id).Delete(g).Error
	return err
}

func (g *Goods) Upload(id int) (info *Goods, err error) {
	var goods Goods
	err = db.Model(g).Where("id = ?", id).Delete(g).Error
	info = &goods
	return info, err
}

func (g *Goods) Get(id int) (info *Goods, err error) {
	var goods Goods
	err = db.Model(g).Where("id = ?", id).First(&goods).Error
	info = &goods
	return info, err
}
