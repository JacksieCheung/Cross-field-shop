package model

import "Cross-field-shop/pkg/constvar"

type CommentsModel struct {
	Id          uint32 `json:"id" gorm:"column:id;not null" binding:"required"`
	UserId      uint32 `json:"user_id" gorm:"column:user_id;not null" binding:"required"`
	CommodityId uint32 `json:"commodity_id" gorm:"column:user_id;not null" binding:"required"`
	Info        string `json:"info" gorm:"column:info" binding:"required"`
	Pictures    string `json:"pictures" gorm:"column:pictures" binding:"required"`
	Tag         string `json:"tag" gorm:"column:tag" binding:"required"`
	Time        string `json:"time" gorm:"column:time" binding:"required"`
	Re          uint8  `json:"re" gorm:"column:re" binding:"required"`
}

type CommentsListItem struct {
	Id          uint32 `json:"id" gorm:"column:id;not null" binding:"required"`
	UserId      uint32 `json:"user_id" gorm:"column:user_id;not null" binding:"required"`
	Name        string `json:"name" gorm:"column:name" binding:"required"`
	Avatar      string `json:"avatar" gorm:"column:avatar" binding:"required"`
	CommodityId uint32 `json:"commodity_id" gorm:"column:user_id;not null" binding:"required"`
	Info        string `json:"info" gorm:"column:info" binding:"required"`
	Pictures    string `json:"pictures" gorm:"column:pictures" binding:"required"`
	Tag         string `json:"tag" gorm:"column:tag" binding:"required"`
	Time        string `json:"time" gorm:"column:time" binding:"required"`
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

func UpdateComment(id, uid uint32, info, tag, pictures string) error {
	var item CommentsModel
	if err := DB.Self.Table("comments").
		Where("id = ? AND user_id = ? AND re = 0", id, uid).
		First(&item).Error; err != nil {
		return err
	}

	item.Info = info
	item.Tag = tag
	item.Pictures = pictures
	err := item.Update()
	if err != nil {
		return err
	}

	return nil
}

func DeleteComment(id, uid, role uint32) error {
	if role > 0 { // 管理员 可以删除任何 评论
		return DB.Self.Table("comment").
			Where("id = ? AND re = 0", id).
			Update("re", 1).Error
	}
	return DB.Self.Table("comment").
		Where("id = ? AND user_id = ? AND re = 0", id, uid).
		Update("re", 1).Error
}

func ListComments(page, limit, uid, commodityId uint32) (*CommentsListItem,
	[]*CommentsListItem, uint64, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	// 计算 offset
	offset := (page - 1) * limit

	list := make([]*CommentsListItem, 0)
	var item *CommentsListItem

	// TODO: tag 要连接到 tagModel 验证其存在
	query := DB.Self.Table("comments").
		Select("comments.*, users.name, users.avatar").
		Joins("left join users on users.id = comments.user_id")

	if err := query.Scan(&list).Where("comments.commodity_id = ?", commodityId).
		Offset(offset).Limit(limit).
		Order("time.id desc").
		Error; err != nil {
		return item, list, uint64(0), err
	}

	if err := query.First(item).Where(
		"comments.commodity_id = ? AND comments.user_id = ?",
		commodityId, uid).
		Error; err != nil {
		return item, list, uint64(0), err
	}

	return item, list, uint64(len(list)), nil
}
