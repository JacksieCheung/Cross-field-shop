package model

type TagsModel struct {
	Id   uint32 `json:"id" gorm:"column:id;not null" binding:"required"`
	Tag  string `json:"tag" gorm:"column:tag" binding:"required"`
	Type uint8  `json:"type" gorm:"column:type" binding:"required"`
	Re   uint8  `json:"re" gorm:"column:re" binding:"required"`
}

// TableName ... 物理表名
func (u *TagsModel) TableName() string {
	return "tags"
}

// Create doc
func (u *TagsModel) Create() error {
	return DB.Self.Create(&u).Error
}

// Update doc
func (u *TagsModel) Update() error {
	return DB.Self.Save(u).Error
}
