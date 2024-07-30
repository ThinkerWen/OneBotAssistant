package course

import (
	"OneBotAssistant/config"
	"OneBotAssistant/util"
	"context"
	"github.com/ThinkerWen/glib-onebot"
	"github.com/ThinkerWen/glib-onebot/api"
	"github.com/ThinkerWen/glib-onebot/events"
	"github.com/ThinkerWen/glib-onebot/messages"
	"github.com/charmbracelet/log"
	"strings"
)

func LoadOnlineCourseEvent(bot *onebot.Bot) {
	loadGroupEvent(bot)
	loadSettingsEvent(bot)
	log.Info("加载 网课搜题助手 成功!")
}

func loadGroupEvent(bot *onebot.Bot) {
	bot.On(events.OnGroupMessageEvent, func(ctx context.Context, event events.IEvent) {
		builder := api.Build(bot.API)
		groupMsg := events.EventParser(event).AsGroupMessage()
		msg := groupMsg.RawMessage
		if !util.IsGroup(config.CONFIG.OnlineCourse.Groups, groupMsg.GroupId) || msg == "" {
			return
		}

		if msg == config.COURSE_HELP_KEY {
			_ = builder.SendGroupMsg(groupMsg.GroupId, false, messages.NewTextMsg(config.COURSE_HELP)).Do(ctx)
			return
		}
		params := strings.Split(msg, " ")
		if config.COURSE_PFX != params[0] || len(params) < 2 {
			return
		}
		result := searchReason(params[1])
		if result == "" {
			_ = builder.SendGroupMsg(groupMsg.GroupId, false, messages.NewTextMsg(config.COURSE_NOT_FOUND)).Do(ctx)
			return
		}
		_ = builder.SendGroupMsg(groupMsg.GroupId, false, messages.NewTextMsg(result)).Do(ctx)
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

		if config.COURSE_ON_KEY == msg {
			util.AddGroup(config.CONFIG.OnlineCourse.Groups, groupMsg.GroupId, "online_course.groups")
			_ = builder.SendGroupMsg(groupMsg.GroupId, false, messages.NewTextMsg(config.COURSE_ON)).Do(ctx)
		} else if config.COURSE_OFF_KEY == msg {
			util.DelGroup(config.CONFIG.OnlineCourse.Groups, groupMsg.GroupId, "online_course.groups")
			_ = builder.SendGroupMsg(groupMsg.GroupId, false, messages.NewTextMsg(config.COURSE_OFF)).Do(ctx)
		}
	})
}
