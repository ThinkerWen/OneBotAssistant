package main

import (
	"OneBotAssistant/config"
	"context"
	"github.com/ThinkerWen/glib-onebot"
	"github.com/charmbracelet/log"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func main() {
	bot, err := onebot.NewBot(config.CONFIG.ApiUrl)
	if err != nil {
		panic(err)
	}

	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Info("Config file changed: " + e.Name)
		_ = viper.Unmarshal(&config.CONFIG)
	})
	viper.WatchConfig()

	LoadAllEvents(bot)
	log.SetLevel(log.InfoLevel)
	err = bot.ListenAndWait(context.Background())
	if err != nil {
		panic(err)
	}
}
