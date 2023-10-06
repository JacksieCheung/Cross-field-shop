package search

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"Data-acquisition-subsystem/handler"
	"Data-acquisition-subsystem/log"
	"Data-acquisition-subsystem/pkg/errno"
	"Data-acquisition-subsystem/service/entity"
	"Data-acquisition-subsystem/service/search"
	"Data-acquisition-subsystem/util"
)

type MhtQueryReq struct {
	StuID string `json:"student_id"`
}

func QueryUserMHT(c *gin.Context) {
	log.Info("QueryUserMHT function called.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	req := &MhtQueryReq{}
	req.StuID = c.Query("student_id")

	mhtQueryRequest := &entity.MhtQueryRequest{}
	_ = util.ConvertEntity(req, mhtQueryRequest)

	service := search.NewQueryUserMhtService(mhtQueryRequest)
	if err := service.QueryUserMht(); err != nil {
		handler.SendBadRequest(c, errno.ErrBind, nil, err.Error(), handler.GetLine())
		return
	}

	handler.SendResponse(c, errno.OK, service.GetResponse())
}
