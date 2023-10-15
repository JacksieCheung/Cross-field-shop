package purchase

type CreatePurchaseReq struct {
	CommodityId uint32 `json:"commodity_id"`
	Number      uint32 `json:"number"`
	Status      uint8  `json:"status"`
}

type UpdatePurchaseReq struct {
	Number uint32 `json:"number"`
}
