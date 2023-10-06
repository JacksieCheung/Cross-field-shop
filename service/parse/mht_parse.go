package parse

import (
	"encoding/json"
	"mime/multipart"
	"regexp"
	"strings"
	"time"

	"go.uber.org/zap"

	"Data-acquisition-subsystem/log"
	"Data-acquisition-subsystem/model"
	"Data-acquisition-subsystem/util"
)

type MhtParser struct {
	BaseParser
	Results [][]string
}

type Chat struct {
	Name    string `json:"name"`
	Time    string `json:"time"`
	Content string `json:"content"`
}

type ChatList struct {
	Chats []Chat `json:"chats"`
}

const mustCompile = `<div style=float:left;margin-right:6px;>(.*?)</div>(.*?)</div><div style=padding-left:20px;><font style="font-size:10pt;font-family:'(.*?)','MS Sans Serif',sans-serif;" color='000000'>(.*?)</font></div>`

func NewMhtParser(file *multipart.FileHeader) *MhtParser {
	return &MhtParser{
		BaseParser: BaseParser{
			File: file,
		},
	}
}

// 解析数据到results中，不同的parser results结构不同
func (m *MhtParser) Parse() {
	m.Data = strings.ReplaceAll(m.Data, "&nbsp;", "")
	reg := regexp.MustCompile(mustCompile)

	results := [][]string{}
	for _, item := range reg.FindAllStringSubmatch(m.Data, -1) {
		results = append(results, item[1:])
	}
	m.Results = results
}

func (m *MhtParser) Process(userID int) error {
	cnt := map[string]int{}
	chats := ChatList{}
	for _, item := range m.Results {
		cnt[item[0]]++
		chats.Chats = append(chats.Chats, Chat{
			Name:    item[0],
			Time:    item[1],
			Content: item[3],
		})
	}

	jsonStr, err := json.Marshal(chats)
	if err != nil {
		return err
	}

	mht := &model.MhtModel{
		Filename:   m.getFilename(),
		UploadTime: time.Now(),
		UploaderID: userID,
		Content:    string(jsonStr),
	}
	if err := mht.Create(); err != nil {
		return err
	}

	total := len(m.Results)
	for name, count := range cnt {
		stuID := util.GetLeftmostDigitsSeq(name)
		participant := &model.ParticipantModel{
			StuID:      stuID,
			MhtID:      mht.ID,
			Proportion: float32(count) / float32(total),
		}

		if err := participant.Create(); err != nil {
			log.Error("Create participant failed",
				zap.String("err", err.Error()))
		}
	}

	return nil
}

func (m *MhtParser) getFilename() string {
	return m.File.Filename
}
