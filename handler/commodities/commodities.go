package commodities

type CreateCommodityReq struct { // 和修改一起用
	Name     string `json:"name"`
	Info     string `json:"info"`
	Price    string `json:"price"`
	Pictures string `json:"pictures"`
	Video    string `json:"video"`
	Remain   int    `json:"remain"`
	Sale     int    `json:"sale"`
	Tag      string `json:"tag"`
}
