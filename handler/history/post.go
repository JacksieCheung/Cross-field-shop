package history

import (
	. "Cross-field-shop/handler"
	"Cross-field-shop/log"
	"Cross-field-shop/model"
	"Cross-field-shop/pkg/errno"
	"Cross-field-shop/util"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Post ... 新增历史，通过 redis 实现
func Post(c *gin.Context) {
	log.Info("history post function called.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	// 从前端获取 商品 id
	var req CreateHistoryReq
	if err := c.ShouldBindJSON(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	// 获取 user_id
	userId := c.MustGet("userID").(uint32)

	// post 商品 id 和 user_id
	history := model.HistoryModel{
		CommodityId: req.CommodityId,
		UserId:      userId,
	}

	// 调用服务
	msg, err := json.Marshal(history)
	if err != nil {
		SendError(c, errno.ErrJsonMarshal, nil, err.Error(), GetLine())
		return
	}

	if err = model.PublishMsg(msg); err != nil {
		log.Error("Publish data error")
		SendError(c, errno.ErrPublishMsg, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, errno.OK, nil)
}
