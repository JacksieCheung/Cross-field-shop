package handler

import (
	"Data-acquisition-subsystem/log"
	"Data-acquisition-subsystem/util"
	"net/http"
	"runtime"
	"strconv"

	"Data-acquisition-subsystem/pkg/errno"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SendResponse(c *gin.Context, err error, data interface{}) {
	code, message := errno.DecodeErr(err)
	log.Info(message,
		zap.String("X-Request-Id", util.GetReqID(c)))

	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

func SendBadRequest(c *gin.Context, err error, data interface{}, cause string, source string) {
	code, message := errno.DecodeErr(err)
	log.Error(message,
		zap.String("X-Request-Id", util.GetReqID(c)),
		zap.String("cause", cause),
		zap.String("source", source))

	c.JSON(http.StatusBadRequest, Response{
		Code:    code,
		Message: message + ": " + cause,
		Data:    data,
	})
}

func SendError(c *gin.Context, err error, data interface{}, cause string, source string) {
	code, message := errno.DecodeErr(err)
	log.Error(message,
		zap.String("X-Request-Id", util.GetReqID(c)),
		zap.String("cause", cause),
		zap.String("source", source))

	c.JSON(http.StatusInternalServerError, Response{
		Code:    code,
		Message: message + ": " + cause,
		Data:    data,
	})
}

func GetLine() string {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		return "muxi-workbench-geteway/handler/handler.go:66"
	}
	return file + ":" + strconv.Itoa(line)
}
