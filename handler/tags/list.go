package tags

import (
	. "Cross-field-shop/handler"
	"Cross-field-shop/log"
	"Cross-field-shop/model"
	"Cross-field-shop/pkg/errno"
	"Cross-field-shop/util"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

// ListTags ... 获取标签
func ListTags(c *gin.Context) {
	log.Info("tags ListTags function called.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	tagType, err := strconv.Atoi(c.Param("type"))
	if err != nil {
		SendBadRequest(c, errno.ErrPathParam, nil, err.Error(), GetLine())
		return
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "20"))
	if err != nil {
		SendBadRequest(c, errno.ErrQuery, nil, err.Error(), GetLine())
		return
	}

	page, err := strconv.Atoi(c.DefaultQuery("page", "0"))
	if err != nil {
		SendBadRequest(c, errno.ErrQuery, nil, err.Error(), GetLine())
		return
	}

	// post 商品 id 和 user_id
	listResp, length, err := model.ListCart(uint32(page), uint32(limit), uint32(tagType))
	if err != nil {
		log.Error(err.Error())
		SendError(c, err, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, errno.OK, gin.H{
		"list": listResp,
		"len":  length,
	})
}
