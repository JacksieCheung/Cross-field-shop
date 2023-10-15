package user

import (
	. "Cross-field-shop/handler"
	"Cross-field-shop/log"
	"Cross-field-shop/pkg/errno"
	"Cross-field-shop/service/user"
	"Cross-field-shop/util"
	"encoding/base64"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Login(c *gin.Context) {
	log.Info("User login function called.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	// 从前端获取 id 和 password
	var req LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	decodePassword, err := base64.StdEncoding.DecodeString(req.Password)
	if err != nil {
		SendError(c, errno.ErrBase64Decode, nil, err.Error(), GetLine())
		return
	}

	// 调用服务
	token, _, err := user.Login(req.Email, string(decodePassword))
	if err != nil {
		SendError(c, errno.ErrPasswordIncorrect, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, errno.OK, &LoginResp{
		Token: token,
	})
}
