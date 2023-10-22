package comments

import (
	. "Cross-field-shop/handler"
	"Cross-field-shop/log"
	"Cross-field-shop/model"
	"Cross-field-shop/pkg/errno"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Post(c *gin.Context) {
	// get request from front end
	var req CreateCommentReq
	if err := c.ShouldBindJSON(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	// 获取 user_id
	userId := c.MustGet("userID").(uint32)

	// 插入数据库
	comment := model.CommentsModel{
		UserId:      userId,
		CommodityId: req.CommodityId,
		Info:        req.Info,
		Pictures:    req.Pictures,
		Tag:         req.Tag,
	}

	err := comment.Create()
	if err != nil {
		log.Error("error while creating comment", zap.String("reason:", err.Error()))
		SendError(c, err, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, errno.OK, nil)
}
