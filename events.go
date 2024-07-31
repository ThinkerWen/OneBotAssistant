package main

import (
	"OneBotAssistant/config"
	"OneBotAssistant/course"
	"OneBotAssistant/hero"
	"OneBotAssistant/molly"
	"OneBotAssistant/reply"
	"OneBotAssistant/sensitive"
	"OneBotAssistant/util"
	"context"
	"github.com/ThinkerWen/glib-onebot"
	"github.com/ThinkerWen/glib-onebot/api"
	"github.com/ThinkerWen/glib-onebot/events"
	"github.com/ThinkerWen/glib-onebot/messages"
	"strconv"
)

func LoadAllEvents(bot *onebot.Bot) {
	loadPluginSettingsEvent(bot)
	if config.CONFIG.HeroPower.Enable {
		hero.LoadHeroPowerEvent(bot)
	}
	if config.CONFIG.OnlineCourse.Enable {
		course.LoadOnlineCourseEvent(bot)
	}
	if config.CONFIG.Molly.Enable {
		molly.LoadMollyEvent(bot)
	}
	if config.CONFIG.Sensitive.Enable {
		sensitive.LoadSensitiveEvent(bot)
	}
	if config.CONFIG.AutoReply.Enable {
		reply.LoadAutoReplyEvent(bot)
	}
}

func loadPluginSettingsEvent(bot *onebot.Bot) {
	bot.On(events.OnGroupMessageEvent, func(ctx context.Context, event events.IEvent) {
		groupMsg := events.EventParser(event).AsGroupMessage()
		if !util.IsHost(groupMsg.Sender.UserId) || groupMsg.RawMessage == "" {
			return
		}
		if len(groupMsg.Message) < 2 || groupMsg.Message[0].Type != "text" {
			return
		}
		prefix := messages.MsgParser(groupMsg.Message[0]).AsText().Text
		switch prefix {
		case config.HOST_ADD_KEY:
			operation(groupMsg.Message, util.AddHost)
		case config.HOST_REMOVE_KEY:
			operation(groupMsg.Message, util.DelHost)
		default:
			return
		}
		_ = api.Build(bot.API).SendGroupMsg(groupMsg.GroupId, false, messages.NewTextMsg(prefix+"成功")).Do(ctx)
	})
}

func operation(message []messages.Message, op func(host int64, key string)) {
	for _, msg := range message {
		if msg.Type != "at" {
			continue
		}
		user := messages.MsgParser(msg).AsAt().QQ
		if uid, err := strconv.ParseInt(user, 10, 64); err == nil {
			op(uid, "hosts")
		}
	}
}
