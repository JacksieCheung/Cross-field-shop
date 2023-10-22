package tags

type CreateTagReq struct {
	Tag  string `json:"tag"`
	Type uint8  `json:"type"`
}

type UpdateTagsReq struct {
	Tag string `json:"tag"`
}
