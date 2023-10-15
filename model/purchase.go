package model

import (
	"Cross-field-shop/pkg/constvar"
)

type PurchaseModel struct {
	Id          uint32 `json:"id" gorm:"column:id;not null" binding:"required"`
	UserId      uint32 `json:"user_id" gorm:"column:user_id;not null" binding:"required"`
	CommodityId uint32 `json:"commodity_id" gorm:"column:commodity_id;not null" binding:"required"`
	Number      uint32 `json:"number" gorm:"column:number" binding:"required"`
	Price       string `json:"price" gorm:"column:price" binding:"required"`
	Status      uint8  `json:"status" gorm:"column:status" binding:"required"`
	Logistics   string `json:"logistics" gorm:"column:logistics" binding:"required"`
	Time        string `json:"time" gorm:"column:time" binding:"required"`
	Re          uint8  `json:"re" gorm:"column:re" binding:"required"`
}

type CartItem struct {
	CommodityId uint32 `json:"commodity_id" gorm:"column:commodity_id;not null" binding:"required"`
	Number      uint32 `json:"number" gorm:"column:number" binding:"required"`
	Price       string `json:"price" gorm:"column:price" binding:"required"`
	Time        string `json:"time" gorm:"column:time" binding:"required"`
	Pictures    string `json:"pictures" gorm:"column:pictures" binding:"required"`
}

// TableName ... 物理表名
func (u *PurchaseModel) TableName() string {
	return "purchase"
}

// Create doc
func (u *PurchaseModel) Create() error {
	return DB.Self.Create(&u).Error
}

// Update doc
func (u *PurchaseModel) Update() error {
	return DB.Self.Save(u).Error
}

func ListCart(page, limit, uid uint32) ([]*CartItem, uint64, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	// 计算 offset
	offset := (page - 1) * limit

	list := make([]*CartItem, 0)

	query := DB.Self.Table("purchase").
		Select("purchase.*, commodities.pictures").
		Joins("left join commodities on purchase.commodity_id = commodities.id").
		Where("purchase.user_id = ? AND purchase.re = 0 AND status = 0", uid).
		Offset(offset).Limit(limit).Order("time.id desc")

	if err := query.Scan(&list).Error; err != nil {
		return list, uint64(0), err
	}

	return list, uint64(len(list)), nil
}

func DeleteCart(id, uid uint32) error {
	return DB.Self.Table("purchase").
		Where("id = ? AND user_id = ? AND re = 0 AND status = 0", id, uid).
		Update("re", 1).Error
}

func UpdateCart(id, uid, num uint32) error {
	var item PurchaseModel
	if err := DB.Self.Table("purchase").
		Where("id = ? AND user_id = ? AND re = 0 AND status = 0", id, uid).
		First(&item).Error; err != nil {
		return err
	}

	item.Number = num
	err := item.Update()
	if err != nil {
		return err
	}

	return nil
}
