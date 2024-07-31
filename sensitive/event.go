package sensitive

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
	"strings"
)

var sensitiveCount = make(map[int64]int)

func LoadSensitiveEvent(core *onebot.Bot) {
	loadGroupEvent(core)
	loadSettingsEvent(core)
	log.Info("加载 敏感词检测 成功!")
}

func loadGroupEvent(bot *onebot.Bot) {
	bot.On(events.OnGroupMessageEvent, func(ctx context.Context, event events.IEvent) {
		groupMsg := events.EventParser(event).AsGroupMessage()
		msg := groupMsg.RawMessage
		if !util.IsGroup(config.CONFIG.Sensitive.Groups, groupMsg.GroupId) || msg == "" {
			return
		}
		if !isForbiddenKeyword(msg, groupMsg.GroupId, 0) && !isSensitive(msg) {
			return
		}

		if _, ok := sensitiveCount[groupMsg.Sender.UserId]; !ok {
			sensitiveCount[groupMsg.Sender.UserId] = 1
		} else if sensitiveCount[groupMsg.Sender.UserId] < config.CONFIG.Sensitive.AlertTimes-1 {
			sensitiveCount[groupMsg.Sender.UserId]++
		} else {
			delete(sensitiveCount, groupMsg.Sender.UserId)
			_ = api.Build(bot.API).SetGroupBan(groupMsg.GroupId, groupMsg.Sender.UserId, config.CONFIG.Sensitive.ShutSeconds).Do(ctx)
		}
		sending := messages.NewTextMsg(fmt.Sprintf("请勿发送不当言论，达到%d次将禁言", config.CONFIG.Sensitive.AlertTimes))
		_ = api.Build(bot.API).SendGroupMsg(groupMsg.GroupId, false, sending).Do(ctx)
		_ = api.Build(bot.API).DeleteMsg(groupMsg.MessageId).Do(ctx)
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

		if config.SENSITIVE_ON_KEY == msg {
			util.AddGroup(config.CONFIG.Sensitive.Groups, groupMsg.GroupId, "sensitive.groups")
			_ = builder.SendGroupMsg(groupMsg.GroupId, false, messages.NewTextMsg(config.SENSITIVE_ON)).Do(ctx)
		} else if config.SENSITIVE_OFF_KEY == msg {
			util.DelGroup(config.CONFIG.Sensitive.Groups, groupMsg.GroupId, "sensitive.groups")
			_ = builder.SendGroupMsg(groupMsg.GroupId, false, messages.NewTextMsg(config.SENSITIVE_OFF)).Do(ctx)
		} else if strings.Contains(msg, config.SENSITIVE_ADD_KEY) {
			params := strings.Split(msg, " ")
			if len(params) != 2 {
				return
			}
			if addSensitiveKeyword(params[1], "*", 0) {
				_ = builder.SendGroupMsg(groupMsg.GroupId, false, messages.NewTextMsg(config.SENSITIVE_ADD_KEY+"成功")).Do(ctx)
			}
		}
	})
}
