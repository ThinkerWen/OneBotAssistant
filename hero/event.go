package hero

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

func LoadHeroPowerEvent(bot *onebot.Bot) {
	loadGroupEvent(bot)
	loadSettingsEvent(bot)
	log.Info("加载 王者荣耀战力查询 成功!")
}

func loadGroupEvent(bot *onebot.Bot) {
	bot.On(events.OnGroupMessageEvent, func(ctx context.Context, event events.IEvent) {
		builder := api.Build(bot.API)
		groupMsg := events.EventParser(event).AsGroupMessage()
		msg := groupMsg.RawMessage
		if !util.IsGroup(config.CONFIG.HeroPower.Groups, groupMsg.GroupId) || msg == "" {
			return
		}

		if msg == config.HERO_HELP_KEY {
			sending := messages.NewTextMsg(config.HERO_HELP)
			_ = builder.SendGroupMsg(groupMsg.GroupId, false, sending).Do(ctx)
			return
		}
		params := strings.Split(msg, " ")
		if config.HERO_PFX != params[0] {
			return
		}
		if len(params) != 3 {
			sending := messages.NewTextMsg(config.HERO_WRONG_TOKEN)
			_ = builder.SendGroupMsg(groupMsg.GroupId, false, sending).Do(ctx)
			return
		}

		hero := params[1]
		server := getHeroServer(params[2])
		if hero == "" || server == "" {
			sending := messages.NewTextMsg(config.HERO_WRONG_TOKEN)
			_ = builder.SendGroupMsg(groupMsg.GroupId, false, sending).Do(ctx)
			return
		}
		sending := messages.NewTextMsg(getHeroPower(hero, server))
		_ = builder.SendGroupMsg(groupMsg.GroupId, false, sending).Do(ctx)
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

		if config.HERO_ON_KEY == msg {
			util.AddGroup(config.CONFIG.HeroPower.Groups, groupMsg.GroupId, "hero_power.groups")
			_ = builder.SendGroupMsg(groupMsg.GroupId, false, messages.NewTextMsg(config.HERO_ON)).Do(ctx)
		} else if config.HERO_OFF_KEY == msg {
			util.DelGroup(config.CONFIG.HeroPower.Groups, groupMsg.GroupId, "hero_power.groups")
			_ = builder.SendGroupMsg(groupMsg.GroupId, false, messages.NewTextMsg(config.HERO_OFF)).Do(ctx)
		}
	})
}
