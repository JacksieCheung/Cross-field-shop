package history

import (
	"Cross-field-shop/log"
	"Cross-field-shop/model"
	"encoding/json"
)

// GoCreateHistory ... 写入 history 数据，子 goroutine 运行
func GoCreateHistory() {
	log.Info("GoCreateHistory Goroutine started.")

	var history = &model.HistoryModel{}

	ch := model.PubSubClient.Self.Channel()
	for msg := range ch {
		log.Info("received")

		if err := json.Unmarshal([]byte(msg.Payload), history); err != nil {
			panic(err)
		}

		if err := history.Create(); err != nil {
			log.Error(err.Error())
		}
	}
}
