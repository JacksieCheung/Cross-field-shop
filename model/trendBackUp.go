package model

type TrendBackUpModel struct {
	Id   uint32 `json:"id" gorm:"column:id;not null" binding:"required"`
	List string `json:"list" gorm:"column:list" binding:"required"`
	Re   uint8  `json:"re" gorm:"column:re" binding:"required"`
}

// TableName ... 物理表名
func (u *TrendBackUpModel) TableName() string {
	return "trend_back_up"
}

// Create doc
func (u *TrendBackUpModel) Create() error {
	return DB.Self.Create(&u).Error
}

// Update doc
func (u *TrendBackUpModel) Update() error {
	return DB.Self.Save(u).Error
}
