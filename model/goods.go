package model

import "gorm.io/gorm"

type Goods struct {
	gorm.Model
	GoodsName  string
	GoodsPrice float64
	GoodsNum   int
}

func (g *Goods) Create(goods Goods) (Goods, error) {
	err := db.Model(g).Create(&goods).Error
	return goods, err
}

func (g *Goods) Delete(id int) error {
	err := db.Model(g).Where("id = ?", id).Delete(g).Error
	return err
}

func (g *Goods) Upload(id int) (Goods, error) {
	var goods Goods
	err := db.Model(g).Where("id = ?", id).Delete(g).Error
	return goods, err
}

func (g *Goods) Get(id int) (Goods, error) {
	var goods Goods
	err := db.Model(g).Where("id = ?", id).First(&goods).Error
	return goods, err
}
