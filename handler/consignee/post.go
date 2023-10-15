package consignee

import (
	. "Cross-field-shop/handler"
	"Cross-field-shop/log"
	"Cross-field-shop/model"
	"Cross-field-shop/pkg/errno"
	"Cross-field-shop/service"
	"Cross-field-shop/util"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Post ... 新增 purchase, 0-加入购物车，1-直接购买，共用同一个 api
func Post(c *gin.Context) {
	log.Info("Consignee post function called.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	// 从前端获取 商品 id
	var req CreateConsigneeReq
	if err := c.ShouldBindJSON(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	// 获取 user_id
	userId := c.MustGet("userID").(uint32)

	// tag 转换成 string
	tagStr := service.ConvertStrListToString(&req.Tag)
	addStr := service.ConvertStrListToString(&req.Address)

	item := model.ConsigneeModel{
		UserId:  userId,
		Address: addStr,
		Name:    req.Name,
		Tel:     req.Tel,
		Tag:     tagStr,
	}

	err := item.Create()
	if err != nil {
		log.Error(err.Error())
		SendError(c, err, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, errno.OK, nil)
}
