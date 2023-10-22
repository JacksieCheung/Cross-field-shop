package tags

import (
	. "Cross-field-shop/handler"
	"Cross-field-shop/log"
	"Cross-field-shop/model"
	"Cross-field-shop/pkg/errno"
	"Cross-field-shop/util"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Post ... 新增 tag
func Post(c *gin.Context) {
	log.Info("tags post function called.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	// 从前端获取 商品 id
	var req CreateTagReq
	if err := c.ShouldBindJSON(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	// post 商品 id 和 user_id
	item := model.TagsModel{
		Tag:  req.Tag,
		Type: req.Type,
	}
	err := item.Create()
	if err != nil {
		log.Error(err.Error())
		SendError(c, err, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, errno.OK, nil)
}
