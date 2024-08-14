package main

import (
	"OneBotAssistant/config"
	"context"
	"github.com/ThinkerWen/glib-onebot"
	"github.com/charmbracelet/log"
)

func main() {
	bot, err := onebot.NewBot(config.CONFIG.ApiUrl)
	if err != nil {
		panic(err)
	}

	LoadAllEvents(bot)
	log.SetLevel(log.InfoLevel)
	err = bot.ListenAndWait(context.Background())
	if err != nil {
		panic(err)
	}
}
