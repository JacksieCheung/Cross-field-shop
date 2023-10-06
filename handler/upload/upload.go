package upload

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"Data-acquisition-subsystem/handler"
	"Data-acquisition-subsystem/log"
	"Data-acquisition-subsystem/pkg/errno"
	"Data-acquisition-subsystem/service/upload"
	"Data-acquisition-subsystem/util"
)

func UploadFile(c *gin.Context) {
	log.Info("UploadFile function called.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	file, err := c.FormFile("file")
	if err != nil {
		handler.SendBadRequest(c, errno.ErrFileNotFound, nil, err.Error(), handler.GetLine())
		return
	}

	uploadSerivce, err := upload.NewUploadService(file, c)
	if err != nil {
		handler.SendError(c, errno.ErrFileInvalid, nil, err.Error(), handler.GetLine())
		return
	}

	if err := uploadSerivce.Upload(); err != nil {
		handler.SendError(c, errno.ErrUploadFailed, nil, err.Error(), handler.GetLine())
		return
	}

	handler.SendResponse(c, errno.OK, "upload file success")
}
