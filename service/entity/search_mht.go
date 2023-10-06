package entity

import (
	"Data-acquisition-subsystem/service/parse"
	"time"
)

type MhtQueryRequest struct {
	StuID string `json:"student_id"`
}

type MhtQueryResponse struct {
	Records []*Record `json:"records"`
}

type Record struct {
	Filename   string         `json:"file_name"`
	Proportion float64        `json:"proportion"`
	Content    parse.ChatList `json:"content"`
	UploadTime time.Time      `json:"upload_time"`
}
