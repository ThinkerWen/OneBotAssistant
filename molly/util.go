package molly

import (
	"OneBotAssistant/config"
	"OneBotAssistant/util"
	"encoding/json"
	"github.com/charmbracelet/log"
	"github.com/go-resty/resty/v2"
	"time"
)

type ContentMolly struct {
	Content  string `json:"content"`
	Type     int    `json:"type"`
	From     string `json:"from"`
	FromName string `json:"fromName"`
	To       string `json:"to"`
	ToName   string `json:"toName"`
}

func mollyChat(content ContentMolly) string {
	headers := make(map[string]string)
	headers["Api-Key"] = config.CONFIG.Molly.ApiKey
	headers["Api-Secret"] = config.CONFIG.Molly.ApiSecret
	headers["Content-Type"] = "application/json;charset=UTF-8"
	client := resty.New()
	client.SetHeaders(headers)
	client.SetTimeout(10 * time.Second)
	if data, errJson := json.Marshal(content); errJson == nil {
		if response, err := util.RequestPOST("https://api.mlyai.com/reply", string(data), headers, client); err == nil {
			return string(response)
		}
	}
	log.Error("Molly 聊天调用失败")
	return ""
}
