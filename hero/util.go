package hero

import (
	"OneBotAssistant/config"
	"OneBotAssistant/util"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/tidwall/gjson"
	"strings"
	"time"
)

func getHeroPower(hero, server string) string {
	api := fmt.Sprintf("https://www.somekey.cn/mini/hero/getHeroInfo.php?hero=%s&type=%s", hero, server)
	client := resty.New()
	client.SetTimeout(10 * time.Second)
	response, err := util.RequestGET(api, nil, client)
	if err != nil || gjson.Get(string(response), "code").Int() != 200 {
		return "请求出错，请联系作者！"
	}
	data := gjson.Get(string(response), "data").String()
	return fmt.Sprintf(config.HERO_POWER_RESULT, gjson.Get(data, "platform").Str, gjson.Get(data, "name").Str, gjson.Get(data, "updatetime").Str,
		gjson.Get(data, "province").Str, gjson.Get(data, "provincePower").Str, gjson.Get(data, "city").Str,
		gjson.Get(data, "cityPower").Str, gjson.Get(data, "area").Str, gjson.Get(data, "areaPower").Str)
}

func getHeroServer(server string) string {
	switch strings.ToUpper(server) {
	case "安卓QQ":
		return "aqq"
	case "安卓微信":
		return "awx"
	case "苹果QQ":
		return "ios_qq"
	case "苹果微信":
		return "ios_wx"
	default:
		return ""
	}
}
