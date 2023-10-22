package model

import "Cross-field-shop/pkg/constvar"

type TagsModel struct {
	Id   uint32 `json:"id" gorm:"column:id;not null" binding:"required"`
	Tag  string `json:"tag" gorm:"column:tag" binding:"required"`
	Type uint8  `json:"type" gorm:"column:type" binding:"required"`
	Re   uint8  `json:"re" gorm:"column:re" binding:"required"`
}

type TagsListItem struct { // by type
	Id  uint32 `json:"id" gorm:"column:id;not null" binding:"required"`
	Tag string `json:"tag" gorm:"column:tag" binding:"required"`
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

func UpdateTags(id uint32, tag string) error {
	var item TagsModel
	if err := DB.Self.Table("tags").
		Where("id = ? AND re = 0", id).
		First(&item).Error; err != nil {
		return err
	}

	item.Tag = tag
	err := item.Update()
	if err != nil {
		return err
	}

	return nil
}

func DeleteTags(id uint32) error {
	return DB.Self.Table("tags").
		Where("id = ? AND re = 0", id).
		Update("re", 1).Error
}

func ListTags(page, limit, tagType uint32) ([]*TagsListItem, uint64, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	// 计算 offset
	offset := (page - 1) * limit

	list := make([]*TagsListItem, 0)

	query := DB.Self.Table("tags").
		Select("tags.*").
		Where("tags.re = 0 AND tags.type = ?", tagType).
		Offset(offset).Limit(limit).Order("time.id desc")

	if err := query.Scan(&list).Error; err != nil {
		return list, uint64(0), err
	}

	return list, uint64(len(list)), nil
}
