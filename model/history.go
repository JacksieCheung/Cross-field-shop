package model

type HistoryModel struct {
	Id          uint32 `json:"id" gorm:"column:id;not null" binding:"required"`
	UserId      uint32 `json:"user_id" gorm:"column:user_id;not null" binding:"required"`
	CommodityId uint32 `json:"commodity_id" gorm:"column:commodity_id;not null" binding:"required"`
	Time        string `json:"time" gorm:"column:time" binding:"required"`
	Re          uint8  `json:"re" gorm:"column:re" binding:"required"`
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
