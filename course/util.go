package course

import (
	"OneBotAssistant/config"
	"OneBotAssistant/util"
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/go-resty/resty/v2"
	"github.com/tidwall/gjson"
	"time"
)

func searchReason(param string) string {
	res := ""
	link := fmt.Sprintf("https://www.hive-net.cn/backend/wangke/search?token=%s&question=%s", config.CONFIG.OnlineCourse.Token, param)
	client := resty.New()
	client.SetTimeout(10 * time.Second)
	response, err := util.RequestGET(link, nil, client)
	if err != nil || gjson.Get(string(response), "code").Int() != 0 {
		log.Error("搜索接口调用失败 Error: ", err)
		return ""
	}
	result := string(response)
	for i, answer := range gjson.Get(result, "data.reasonList").Array() {
		question := answer.Get("question").Str
		options := answer.Get("options").Str
		reason := answer.Get("reason").Str
		data := fmt.Sprintf("问题:\n%s", question)
		if i != 0 {
			data = "================\n" + data
		}
		if options != "" && options != "无" {
			data += fmt.Sprintf("\n选项:\n%s", options)
		}
		data += fmt.Sprintf("\n答案:\n%s\n", reason)
		res += data
		if i == config.CONFIG.OnlineCourse.Limit-1 {
			return res
		}
	}
	return res
}
