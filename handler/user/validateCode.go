package user

import (
	. "Cross-field-shop/handler"
	"Cross-field-shop/log"
	"Cross-field-shop/model"
	"Cross-field-shop/pkg/errno"
	"Cross-field-shop/util"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"math/rand"
	"strings"
	"time"
)

// ValidateCode ... 生成验证码存到 redis，然后发送邮件出去
func ValidateCode(c *gin.Context) {
	// 生成验证码
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < 6; i++ {
		_, err := fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
		if err != nil {
			log.Error("generate code error",
				zap.String("reason:", err.Error()))
			SendError(c, err, nil, err.Error(), GetLine())
			return
		}
	}
	code := sb.String()

	// 从前端 Email
	var req EmailValidateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	// push code to redis, format: mail:code, overtime: 2mins
	ifHas, err := model.HasExistedInRedis(req.Email)
	if err != nil {
		log.Error("redis call error", zap.String("reason:", err.Error()))
		SendError(c, err, nil, err.Error(), GetLine())
		return
	}
	if ifHas { // 已经存在
		log.Info("mail is already existed, try after 2mins")
		SendBadRequest(c, errors.New("too often call, try after 2mins"),
			nil, "", GetLine())
		return
	}

	err = model.SetStringInRedis(req.Email, code, 2*time.Minute)
	if err != nil {
		log.Error("redis call error", zap.String("reason:", err.Error()))
		SendError(c, err, nil, err.Error(), GetLine())
		return
	}

	// send email
	err = util.SendEmail(req.Email, code)
	if err != nil {
		log.Error("send email error", zap.String("reason:", err.Error()))
		SendError(c, err, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, errno.OK, nil)
}
