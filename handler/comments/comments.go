package comments

type CreateCommentReq struct {
	CommodityId uint32 `json:"commodity_id"`
	Info        string `json:"info"`
	Pictures    string `json:"pictures"` // TODO: upload 七牛云
	Tag         string `json:"tag"`
}

type UpdateCommentReq struct {
	Info     string `json:"info"`
	Pictures string `json:"pictures"` // TODO: upload 七牛云
	Tag      string `json:"tag"`
}
