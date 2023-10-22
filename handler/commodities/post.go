package commodities

import (
	. "Cross-field-shop/handler"
	"Cross-field-shop/log"
	"Cross-field-shop/model"
	"Cross-field-shop/pkg/errno"
	"Cross-field-shop/util"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Post ... 新增 commodity
func Post(c *gin.Context) {
	log.Info("Commodities post function called.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	// 从前端获取 商品 id
	var req CreateCommodityReq
	if err := c.ShouldBindJSON(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	// post 商品
	commodity := model.CommoditiesModel{
		Name:     req.Name,
		Pictures: req.Pictures,
		Info:     req.Info,
		Price:    req.Price,
		Tag:      req.Tag,
		Sale:     req.Sale,
		Video:    req.Video,
		Remain:   req.Remain,
	}
	err := commodity.Create()
	if err != nil {
		log.Error(err.Error())
		SendError(c, err, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, errno.OK, nil)
}
