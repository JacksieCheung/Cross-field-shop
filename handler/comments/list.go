package comments

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

// List ... 获取评论列表 by commodity，同时获取我的评论（如果有）
func List(c *gin.Context) {
	log.Info("Comments List function called.",
		zap.String("X-Request-Id", util.GetReqID(c)))

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

	// 获取 user_id
	userId := c.MustGet("userID").(uint32)

	commodityId, err := strconv.Atoi(c.Param("commodity_id"))
	if err != nil {
		SendBadRequest(c, err, nil, err.Error(), GetLine())
		return
	}

	// list comments
	item, listResp, length, err := model.ListComments(uint32(page), uint32(limit),
		userId, uint32(commodityId))
	if err != nil {
		log.Error(err.Error())
		SendError(c, err, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, errno.OK, gin.H{
		"list":          listResp,
		"owner_comment": item,
		"len":           length,
	})
}
