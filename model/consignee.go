package model

type ConsigneeModel struct {
	Id      uint32 `json:"id" gorm:"column:id;not null" binding:"required"`
	UserId  uint32 `json:"user_id" gorm:"column:user_id;not null" binding:"required"`
	Address string `json:"address" gorm:"column:address" binding:"required"`
	Name    string `json:"name" gorm:"column:name" binding:"required"`
	Tel     string `json:"tel" gorm:"column:tel" binding:"required"`
	Tag     string `json:"tag" gorm:"column:tag" binding:"required"`
	Re      uint8  `json:"re" gorm:"column:re" binding:"required"`
}

// TableName ... 物理表名
func (u *ConsigneeModel) TableName() string {
	return "consignee"
}

// Create doc
func (u *ConsigneeModel) Create() error {
	return DB.Self.Create(&u).Error
}

// Update doc
func (u *ConsigneeModel) Update() error {
	return DB.Self.Save(u).Error
}
