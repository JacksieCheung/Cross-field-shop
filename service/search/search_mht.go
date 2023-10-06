package search

import (
	"Data-acquisition-subsystem/model"
	"Data-acquisition-subsystem/service/entity"
	"Data-acquisition-subsystem/service/parse"
	"encoding/json"
)

type QueryUserMhtService struct {
	req  *entity.MhtQueryRequest
	resp *entity.MhtQueryResponse
}

func NewQueryUserMhtService(request *entity.MhtQueryRequest) *QueryUserMhtService {
	return &QueryUserMhtService{
		req:  request,
		resp: new(entity.MhtQueryResponse),
	}
}

func (qs *QueryUserMhtService) QueryUserMht() error {
	participants, err := model.GetMhtsByStuID(qs.req.StuID)
	if err != nil {
		return err
	}

	for _, p := range participants {
		mht, err := model.GetMhtByID(p.MhtID)
		if err != nil {
			return err
		}
		chatList := parse.ChatList{}
		err = json.Unmarshal([]byte(mht.Content), &chatList)
		if err != nil {
			return err
		}
		qs.resp.Records = append(qs.resp.Records, &entity.Record{
			Filename:   mht.Filename,
			Proportion: float64(p.Proportion),
			Content:    chatList,
			UploadTime: mht.UploadTime,
		})
	}

	return nil
}

func (qs *QueryUserMhtService) GetResponse() *entity.MhtQueryResponse {
	return qs.resp
}
