package molly

import (
	"OneBotAssistant/config"
	"OneBotAssistant/util"
	"context"
	"fmt"
	"github.com/ThinkerWen/glib-onebot"
	"github.com/ThinkerWen/glib-onebot/api"
	"github.com/ThinkerWen/glib-onebot/events"
	"github.com/ThinkerWen/glib-onebot/messages"
	"github.com/charmbracelet/log"
	"github.com/tidwall/gjson"
	"regexp"
	"strconv"
)

func LoadMollyEvent(bot *onebot.Bot) {
	loadGroupEvent(bot)
	loadSettingsEvent(bot)
	log.Info("加载 莫莉云机器人 成功!")
}

func loadGroupEvent(bot *onebot.Bot) {
	bot.On(events.OnGroupMessageEvent, func(ctx context.Context, event events.IEvent) {
		builder := api.Build(bot.API)
		groupMsg := events.EventParser(event).AsGroupMessage()
		msg := regexp.MustCompile(`\[CQ[^]]*]`).ReplaceAllString(groupMsg.RawMessage, "")
		if !util.IsGroup(config.CONFIG.Molly.Groups, groupMsg.GroupId) || msg == "" {
			return
		}
		for i, m := range groupMsg.Message {
			if m.Type == "at" {
				if uid, err := strconv.ParseInt(messages.MsgParser(m).AsAt().QQ, 10, 64); err == nil {
					if uid == config.CONFIG.Molly.QQ {
						break
					}
				}
			}
			if i == len(groupMsg.Message)-1 {
				return
			}
		}
		data := new(ContentMolly)
		data.Type = 2
		data.Content = msg
		data.To = strconv.FormatInt(groupMsg.GroupId, 10)
		data.From = strconv.FormatInt(groupMsg.Sender.UserId, 10)
		result := mollyChat(*data)
		if result == "" || gjson.Get(result, "code").Str != "00000" || len(gjson.Get(result, "data").Array()) == 0 {
			sending1 := messages.NewTextMsg("我现在不想说话")
			sending2 := messages.NewFaceMsg("179")
			_ = builder.SendGroupMsg(groupMsg.GroupId, false, sending1, sending2).Do(ctx)
			return
		}
		_ = builder.SendGroupMsg(groupMsg.GroupId, false, messages.NewTextMsg(gjson.Get(result, "data|0.content").Str)).Do(ctx)
	})
}

func loadSettingsEvent(bot *onebot.Bot) {
	bot.On(events.OnGroupMessageEvent, func(ctx context.Context, event events.IEvent) {
		builder := api.Build(bot.API)
		groupMsg := events.EventParser(event).AsGroupMessage()
		msg := groupMsg.RawMessage
		if !util.IsHost(groupMsg.Sender.UserId) || msg == "" {
			return
		}

		if config.MOLLY_ON_KEY == msg {
			util.AddGroup(config.CONFIG.Molly.Groups, groupMsg.GroupId, "molly.groups")
			sending := fmt.Sprintf(config.MOLLY_ON, config.CONFIG.Molly.Name)
			_ = builder.SendGroupMsg(groupMsg.GroupId, false, messages.NewTextMsg(sending)).Do(ctx)
		} else if config.MOLLY_OFF_KEY == msg {
			util.DelGroup(config.CONFIG.Molly.Groups, groupMsg.GroupId, "molly.groups")
			sending := fmt.Sprintf(config.MOLLY_OFF, config.CONFIG.Molly.Name)
			_ = builder.SendGroupMsg(groupMsg.GroupId, false, messages.NewTextMsg(sending)).Do(ctx)
		}
	})
}
