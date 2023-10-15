package commodities

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

// List ... 获取商品列表/我的商品
func List(c *gin.Context) {
	log.Info("commodities List function called.",
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

	ifUser, err := strconv.Atoi(c.DefaultQuery("user", "0"))
	if err != nil {
		SendBadRequest(c, errno.ErrQuery, nil, err.Error(), GetLine())
		return
	}

	// 获取 user_id
	userId := c.MustGet("userID").(uint32)

	// post 商品 id 和 user_id
	listResp, length, err := model.ListCommodities(uint32(page), uint32(limit), userId, uint32(ifUser))
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
