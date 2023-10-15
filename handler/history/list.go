package history

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

// List ... 前端做成 yj app 那种一列一列下来的 界面
func List(c *gin.Context) {
	log.Info("history list function called.",
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

	// 获取 uid
	userID := c.MustGet("userID").(uint32)

	// 构造 list 请求
	listResp, length, err := model.ListHistory(uint32(page), uint32(limit), userID)
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, errno.OK, gin.H{
		"list": listResp,
		"len":  length,
	})
}
