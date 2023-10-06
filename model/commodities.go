package model

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
