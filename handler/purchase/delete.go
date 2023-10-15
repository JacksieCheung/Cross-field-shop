package purchase

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

// DeleteCart ... 删除购物车
func DeleteCart(c *gin.Context) {
	log.Info("purchase DeleteCart function called.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		SendBadRequest(c, errno.ErrPathParam, nil, err.Error(), GetLine())
		return
	}

	// 获取 user_id
	userId := c.MustGet("userID").(uint32)

	// post 商品 id 和 user_id
	err = model.DeleteCart(uint32(id), userId)
	if err != nil {
		log.Error(err.Error())
		SendError(c, err, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, errno.OK, nil)
}
