package model

import "Cross-field-shop/pkg/constvar"

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

func ListConsignee(page, limit, uid uint32) ([]*ConsigneeModel, uint64, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	// 计算 offset
	offset := (page - 1) * limit

	list := make([]*ConsigneeModel, 0)

	query := DB.Self.Table("consignee").
		Where("consignee.user_id = ? AND purchase.re = 0", uid).
		Offset(offset).Limit(limit).Order("time.id desc")

	if err := query.Scan(&list).Error; err != nil {
		return list, uint64(0), err
	}

	return list, uint64(len(list)), nil
}

func UpdateConsignee(id, uid uint32, address, tags, name, tel string) error {
	var item ConsigneeModel
	if err := DB.Self.Table("consignee").
		Where("id = ? AND user_id = ? AND re = 0", id, uid).
		First(&item).Error; err != nil {
		return err
	}

	if address != "" {
		item.Address = address
	}
	if tags != "" {
		item.Tag = tags
	}
	if name != "" {
		item.Name = name
	}
	if tel != "" {
		item.Tel = tel
	}

	err := item.Update()
	if err != nil {
		return err
	}

	return nil
}

func DeleteConsignee(id, uid uint32) error {
	return DB.Self.Table("consignee").
		Where("id = ? AND user_id = ? AND re = 0", id, uid).
		Update("re", 1).Error
}
