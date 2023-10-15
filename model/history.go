package model

import "Cross-field-shop/pkg/constvar"

type HistoryModel struct {
	Id          uint32 `json:"id" gorm:"column:id;not null" binding:"required"`
	UserId      uint32 `json:"user_id" gorm:"column:user_id;not null" binding:"required"`
	CommodityId uint32 `json:"commodity_id" gorm:"column:commodity_id;not null" binding:"required"`
	Time        string `json:"time" gorm:"column:time" binding:"required"`
	Re          uint8  `json:"re" gorm:"column:re" binding:"required"`
}

type HistoryListItem struct {
	CommodityId uint32 `json:"commodity_id" gorm:"column:commodity_id;not null" binding:"required"`
	Time        string `json:"time" gorm:"column:time" binding:"required"`
	Name        string `json:"name" gorm:"column:name" binding:"required"`
	Price       string `json:"price" gorm:"column:price" binding:"required"`
	Pictures    string `json:"pictures" gorm:"column:pictures" binding:"required"`
}

// TableName ... 物理表名
func (u *HistoryModel) TableName() string {
	return "history"
}

// Create doc
func (u *HistoryModel) Create() error {
	return DB.Self.Create(&u).Error
}

// Update doc
func (u *HistoryModel) Update() error {
	return DB.Self.Save(u).Error
}

// ListHistory ... 分页
func ListHistory(page, limit, uid uint32) ([]*HistoryListItem, uint64, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	// 计算 offset
	offset := (page - 1) * limit

	list := make([]*HistoryListItem, 0)

	query := DB.Self.Table("history").
		Select("history.*, users.name, users.avatar, users.group_id").
		Where("history.user_id = ? AND history.re = 0", uid).
		Joins("left join commodities on history.commodity_id = commodities.id").
		Offset(offset).Limit(limit).Order("time.id desc")

	if err := query.Scan(&list).Error; err != nil {
		return list, uint64(0), err
	}

	return list, uint64(len(list)), nil
}
