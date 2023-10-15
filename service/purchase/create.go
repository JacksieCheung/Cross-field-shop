package purchase

import (
	"Cross-field-shop/model"
	"errors"
	"github.com/jinzhu/gorm"
)

func Create(uid, commodityId, num uint32, status uint8) error {
	// 根据 CommodityId 去抓 commodity 的 price
	item, err := model.GetCommodityById(commodityId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("no commodity")
		}
		return err
	}

	p := model.PurchaseModel{
		UserId:      uid,
		CommodityId: commodityId,
		Number:      num,
		Price:       item.Price,
		Status:      status,
	}

	if err = p.Create(); err != nil {
		return err
	}

	return nil
}
