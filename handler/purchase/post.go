package purchase

import (
	. "Cross-field-shop/handler"
	"Cross-field-shop/log"
	"Cross-field-shop/pkg/errno"
	"Cross-field-shop/service/purchase"
	"Cross-field-shop/util"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Post ... 新增 purchase, 0-加入购物车，1-直接购买，共用同一个 api
func Post(c *gin.Context) {
	log.Info("purchase post function called.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	// 从前端获取 商品 id
	var req CreatePurchaseReq
	if err := c.ShouldBindJSON(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	// 获取 user_id
	userId := c.MustGet("userID").(uint32)

	// post 商品 id 和 user_id
	err := purchase.Create(userId, req.CommodityId, req.Number, req.Status)
	if err != nil {
		log.Error(err.Error())
		SendError(c, err, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, errno.OK, nil)
}
