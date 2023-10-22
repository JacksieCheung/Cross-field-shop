package user

import (
	. "Cross-field-shop/handler"
	"Cross-field-shop/log"
	"Cross-field-shop/model"
	"Cross-field-shop/pkg/errno"
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Register ... must call ValidateCode before
func Register(c *gin.Context) {
	// get request from front end
	var req RegisterReq
	if err := c.ShouldBindJSON(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	// validate code from redis
	code, ifHas, err := model.GetStringFromRedis(req.Email)
	if err != nil {
		log.Error("error whild redis call", zap.String("reason:", err.Error()))
		SendError(c, err, nil, err.Error(), GetLine())
		return
	}
	if !ifHas { // 找不到
		log.Info("code maybe expired, try again")
		SendBadRequest(c, errors.New("code maybe expired, try again"),
			nil, "", GetLine())
		return
	}

	if code != req.ValidateCode {
		log.Info("invalid validate code")
		SendBadRequest(c, errors.New("invalid validate code"), nil, "", GetLine())
		return
	}

	// 验证通过，插入数据库
	user := model.UserModel{
		Email:    req.Email,
		Password: req.Password,
		Name:     req.Name,
	}

	err = user.Create()
	if err != nil {
		log.Error("error while creating user", zap.String("reason:", err.Error()))
		SendError(c, err, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, errno.OK, nil)
}
