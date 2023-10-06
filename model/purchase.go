package model

type PurchaseModel struct {
	Id          uint32 `json:"id" gorm:"column:id;not null" binding:"required"`
	UserId      uint32 `json:"user_id" gorm:"column:user_id;not null" binding:"required"`
	CommodityId uint32 `json:"commodity_id" gorm:"column:commodity_id;not null" binding:"required"`
	Number      int    `json:"number" gorm:"column:number" binding:"required"`
	Price       string `json:"price" gorm:"column:price" binding:"required"`
	Status      uint8  `json:"status" gorm:"column:status" binding:"required"`
	Logistics   string `json:"logistics" gorm:"column:logistics" binding:"required"`
	Time        string `json:"time" gorm:"column:time" binding:"required"`
	Re          uint8  `json:"re" gorm:"column:re" binding:"required"`
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
