package upload

import (
	"mime/multipart"

	"Data-acquisition-subsystem/log"
	"Data-acquisition-subsystem/service/parse"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UploadService struct {
	parser parse.Parser
	ctx    *gin.Context
}

func NewUploadService(file *multipart.FileHeader, ctx *gin.Context) (*UploadService, error) {
	parse, err := parse.ParserFactory(file)
	if err != nil {
		return nil, err
	}
	return &UploadService{parse, ctx}, nil
}

func (us *UploadService) Upload() error {
	userID := us.ctx.MustGet("userID").(int)
	if err := us.parser.ReadFile(); err != nil {
		log.Error("error happend in readfile",
			zap.String("error", err.Error()))
		return err
	}

	us.parser.Parse()
	if err := us.parser.Process(userID); err != nil {
		log.Error("error happend in processing",
			zap.String("error", err.Error()))
		return err
	}

	if err := us.parser.SaveFile(); err != nil {
		log.Error("error happend in savefile",
			zap.String("error", err.Error()))
		return err
	}

	return nil
}
