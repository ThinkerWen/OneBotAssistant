package reply

import (
	"OneBotAssistant/config"
	"OneBotAssistant/util"
	"context"
	"encoding/json"
	"github.com/ThinkerWen/glib-onebot"
	"github.com/ThinkerWen/glib-onebot/api"
	"github.com/ThinkerWen/glib-onebot/events"
	"github.com/ThinkerWen/glib-onebot/messages"
	"github.com/charmbracelet/log"
	"strconv"
)

var sequence = make([]Sequence, 0)

func LoadAutoReplyEvent(core *onebot.Bot) {
	loadGroupEvent(core)
	loadSettingsEvent(core)
	log.Info("加载 自动回复助手 成功!")
}

func loadGroupEvent(bot *onebot.Bot) {
	bot.On(events.OnGroupMessageEvent, func(ctx context.Context, event events.IEvent) {
		builder := api.Build(bot.API)
		groupMsg := events.EventParser(event).AsGroupMessage()
		msg := groupMsg.RawMessage
		if !util.IsGroup(config.CONFIG.AutoReply.Groups, groupMsg.GroupId) {
			return
		}

		if answer := matchAsk(msg, groupMsg.GroupId, 0); answer != "" {
			var message []messages.Message
			if err := json.Unmarshal([]byte(answer), &message); err == nil {
				_ = builder.SendGroupMsg(groupMsg.GroupId, false, message...).Do(ctx)
			}
			return
		}

		if msg == config.REPLY_ADD {
			sequence = append(sequence, Sequence{Sender: groupMsg.Sender.UserId, Receiver: groupMsg.GroupId, State: AutoReplyState{Step: 0, Range: "*", RangeType: 0}})
			_ = builder.SendGroupMsg(groupMsg.GroupId, false, messages.NewTextMsg(config.REPLY_ASK)).Do(ctx)
			return
		}

		for i, v := range sequence {
			if !(v.Sender == groupMsg.Sender.UserId && v.Receiver == groupMsg.GroupId) {
				continue
			}
			message, _ := json.Marshal(groupMsg.Message)
			currentState := checkReply(msg, string(message), &v.State)
			switch currentState {
			case 1:
				_ = builder.SendGroupMsg(groupMsg.GroupId, false, messages.NewTextMsg(config.REPLY_ANSWER)).Do(ctx)
			case 2:
				_ = builder.SendGroupMsg(groupMsg.GroupId, false, messages.NewTextMsg(config.REPLY_RANGE)).Do(ctx)
			case 3:
				if msg == "1" {
					v.State.Range = strconv.FormatInt(groupMsg.GroupId, 10)
				}
				saveReply(v)
				_ = builder.SendGroupMsg(groupMsg.GroupId, false, messages.NewTextMsg(config.REPLY_ADD_SUCCESS)).Do(ctx)
				sequence = removeSequence(groupMsg.Sender.UserId, groupMsg.GroupId)
				return
			default:
				return
			}
			sequence[i] = v
			return
		}
	})
}

func loadSettingsEvent(bot *onebot.Bot) {
	bot.On(events.OnGroupMessageEvent, func(ctx context.Context, event events.IEvent) {
		groupMsg := events.EventParser(event).AsGroupMessage()
		msg := groupMsg.RawMessage
		if !util.IsHost(groupMsg.Sender.UserId) || msg == "" {
			return
		}

		if config.REPLY_ON_KEY == msg {
			util.AddGroup(config.CONFIG.AutoReply.Groups, groupMsg.GroupId, "auto_reply.groups")
		} else if config.REPLY_OFF_KEY == msg {
			util.DelGroup(config.CONFIG.AutoReply.Groups, groupMsg.GroupId, "auto_reply.groups")
		} else {
			return
		}
		_ = api.Build(bot.API).SendGroupMsg(groupMsg.GroupId, false, messages.NewTextMsg(msg+"成功")).Do(ctx)
	})
}
