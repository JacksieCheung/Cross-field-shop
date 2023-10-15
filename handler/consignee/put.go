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
	"strconv"
)

// UpdateConsignee ... 修改地址 --- address/name/tag/tel
func UpdateConsignee(c *gin.Context) {
	log.Info("consignee UpdateConsignee function called.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		SendBadRequest(c, errno.ErrPathParam, nil, err.Error(), GetLine())
		return
	}

	// 获取 user_id
	userId := c.MustGet("userID").(uint32)

	var req UpdateConsigneeReq
	if err := c.ShouldBindJSON(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	// update
	err = model.UpdateConsignee(uint32(id), userId,
		service.ConvertStrListToString(&req.Address),
		service.ConvertStrListToString(&req.Tag), req.Name, req.Tel)
	if err != nil {
		log.Error(err.Error())
		SendError(c, err, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, errno.OK, nil)
}
