package model

//type Commodity struct {
//	gorm.Model
//	GoodsName   string `gorm:"type:varchar(50)"`  //商品名称
//	Description string `gorm:"type:varchar(255)"` //商品描述
//	GoodsStock  int64  `gorm:"type:int(5)"`       //商品库存
//}
//
//type Category struct {
//	gorm.Model
//	GoodsCategory int64  `gorm:"type:int(3)"`      //商品第一分类id
//	CategoryName  string `gorm:"type:varchar(50)"` //分类名称
//}
//
//type Secondary struct {
//	gorm.Model
//	CategoryID            int64   `gorm:"type:int(3)"`        //商品第一分类id
//	SecondaryCategoryName string  `gorm:"type:varchar(50)"`   //副分类名称
//	Price                 float64 `gorm:"type:decimal(10,2)"` //商品分类价格
//}

//type FromHttp struct {
//	Commodity Commodity           `json:"goods"`
//	Price     string              `json:"price"`
//	Stock     string              `json:"stock"`
//	Property  map[string][]string `json:"property"`
//}

// 商品表
//type Commodity struct {
//	gorm.Model
//	GoodsName   string `gorm:"index"`             //商品名
//	Image       string `gorm:"type:text(0)"`      //头图
//	Description string `gorm:"type:varchar(250)"` //描述
//	Brand       int64  `gorm:"type:varchar(100)"` //品牌
//	Status      int64  `gorm:"type:tinyint(1)"`   //状态
//	Supplier    string `gorm:"type:varchar(255)"` //供应商
//}
//
//// 商品属性表
//type Property struct {
//	gorm.Model
//	GoodsId int64  `gorm:"index,type:int(10)"` //商品id
//	Name    string `gorm:"type:varchar(100)"`  //商品属性名
//}
//
//// 商品属性
//type PropertyValue struct {
//	gorm.Model
//	PropertyId int64  `gorm:"index,type:int(10)"` //属性id
//	Value      string `gorm:"type:varchar(100)"`  //属性值
//}
//
//// 商品sku
//type Sku struct {
//	gorm.Model
//	GoodsId     int64  `gorm:"index,type:int(10)"` //商品id
//	ProValIdFir int64  `gorm:"index,type:int(10)"` //第一个商品属性值
//	ProValIdSec int64  `gorm:"index,type:int(10)"` //第二个商品属性值
//	ProValIdThi int64  `gorm:"index,type:int(10)"` //第三个商品属性值
//	Price       string `gorm:"type:decimal(10,2)"` //单价
//	Stock       int64  `gorm:"index,type:int(6)"`  //库存
//}
