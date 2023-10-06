package model

type CommentsModel struct {
	Id       uint32 `json:"id" gorm:"column:id;not null" binding:"required"`
	Info     string `json:"info" gorm:"column:info" binding:"required"`
	Pictures string `json:"pictures" gorm:"column:pictures" binding:"required"`
	Tag      string `json:"tag" gorm:"column:tag" binding:"required"`
	Time     string `json:"time" gorm:"column:time" binding:"required"`
	Re       uint8  `json:"re" gorm:"column:re" binding:"required"`
}

// TableName ... 物理表名
func (u *CommentsModel) TableName() string {
	return "comments"
}

// Create doc
func (u *CommentsModel) Create() error {
	return DB.Self.Create(&u).Error
}

// Update doc
func (u *CommentsModel) Update() error {
	return DB.Self.Save(u).Error
}
