package consignee

type CreateConsigneeReq struct {
	Address []string `json:"address"`
	Name    string   `json:"name"`
	Tel     string   `json:"tel"`
	Tag     []string `json:"tag"` // 应该获得一个 list，转成 string
}

type UpdateConsigneeReq struct {
	Address []string `json:"address"`
	Name    string   `json:"name"`
	Tel     string   `json:"tel"`
	Tag     []string `json:"tag"`
}
