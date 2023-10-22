package model

import (
	"Cross-field-shop/handler/commodities"
	"Cross-field-shop/pkg/constvar"
)

type CommoditiesModel struct {
	Id       uint32 `json:"id" gorm:"column:id;not null" binding:"required"`
	Name     string `json:"name" gorm:"column:name" binding:"required"`
	Info     string `json:"info" gorm:"column:info" binding:"required"`
	Price    string `json:"price" gorm:"column:price" binding:"required"`
	Pictures string `json:"pictures" gorm:"column:pictures" binding:"required"`
	Video    string `json:"video" gorm:"column:video" binding:"required"`
	Remain   int    `json:"remain" gorm:"column:remain" binding:"required"`
	Sale     int    `json:"sale" gorm:"column:sale" binding:"required"`
	Tag      string `json:"tag" gorm:"column:tag" binding:"required"`
	Re       uint8  `json:"re" gorm:"column:re" binding:"required"`
}

// TableName ... 物理表名
func (u *CommoditiesModel) TableName() string {
	return "commodities"
}

// Create doc
func (u *CommoditiesModel) Create() error {
	return DB.Self.Create(&u).Error
}

// Update doc
func (u *CommoditiesModel) Update() error {
	return DB.Self.Save(u).Error
}

func GetCommodityById(id uint32) (*CommoditiesModel, error) {
	u := &CommoditiesModel{}
	d := DB.Self.Table("commodities").Where("id = ?", id).First(u)
	return u, d.Error
}

func UpdateCommodity(id uint32, req *commodities.CreateCommodityReq) error {
	var item CommoditiesModel
	if err := DB.Self.Table("commodities").
		Where("id = ? AND AND re = 0", id).
		First(&item).Error; err != nil {
		return err
	}

	item.Name = req.Name
	item.Sale = req.Sale
	item.Tag = req.Tag
	item.Video = req.Video
	item.Remain = req.Remain
	item.Price = req.Price
	item.Info = req.Info
	item.Pictures = req.Pictures
	err := item.Update()
	if err != nil {
		return err
	}

	return nil
}

func DeleteCommodity(id uint32) error {
	return DB.Self.Table("commodities").
		Where("id = ? AND re = 0", id).
		Update("re", 1).Error
}

func ListCommodities(page, limit, uid, ifUser, ifSale uint32) ([]*CommoditiesModel, uint64, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	// 计算 offset
	offset := (page - 1) * limit

	list := make([]*CommoditiesModel, 0)

	// TODO: tag 要连接到 tagModel 验证其存在
	query := DB.Self.Table("Commodities").
		Select("purchase.*, commodities.pictures").
		Joins("left join commodities on purchase.commodity_id = commodities.id")

	if ifUser != 0 {
		query = query.Where("purchase.user_id = ?", uid)
	}

	if ifSale != 0 {
		query = query.Order("commodities.sale desc")
	}

	if err := query.Scan(&list).Where("commodities.re = 0").
		Offset(offset).Limit(limit).
		Order("time.id desc").
		Error; err != nil {
		return list, uint64(0), err
	}

	return list, uint64(len(list)), nil
}
