package util

import (
	"OneBotAssistant/config"
	"github.com/charmbracelet/log"
	"github.com/spf13/viper"
	"strings"
)

func IsHost(host int64) bool {
	for _, v := range config.CONFIG.Hosts {
		if v == host {
			return true
		}
	}
	return false
}

func IsGroup(groups []int64, group int64) bool {
	for _, v := range groups {
		if v == group {
			return true
		}
	}
	return false
}

func AddHost(host int64, key string) {
	if IsHost(host) {
		return
	}
	config.CONFIG.Hosts = append(config.CONFIG.Hosts, host)
	viper.Set(key, config.CONFIG.Hosts)
	if err := config.SaveConfig(); err != nil {
		log.Error("Error writing config: ", err)
	}
}

func DelHost(host int64, key string) {
	var result []int64
	for _, v := range config.CONFIG.Hosts {
		if v != host {
			result = append(result, v)
		}
	}
	viper.Set(key, result)
	if err := config.SaveConfig(); err != nil {
		log.Error("Error writing config: ", err)
	}
}

func AddGroup(groups []int64, group int64, key string) {
	if IsGroup(groups, group) {
		return
	}
	newGroups := append(groups[:0:0], groups...)
	newGroups = append(newGroups, group)
	setConfig(newGroups, strings.Split(key, ".")[0])
	if err := config.SaveConfig(); err != nil {
		log.Error("Error writing config: ", err)
	}
}

func DelGroup(groups []int64, group int64, key string) {
	var result []int64
	for _, v := range groups {
		if v != group {
			result = append(result, v)
		}
	}
	setConfig(result, strings.Split(key, ".")[0])
	if err := config.SaveConfig(); err != nil {
		log.Error("Error writing config: ", err)
	}
}

func setConfig(groups []int64, key string) {
	switch key {
	case "molly":
		config.CONFIG.Molly.Groups = groups
		viper.Set(key, config.CONFIG.Molly)
	case "hero_power":
		config.CONFIG.HeroPower.Groups = groups
		viper.Set(key, config.CONFIG.HeroPower)
	case "sensitive":
		config.CONFIG.Sensitive.Groups = groups
		viper.Set(key, config.CONFIG.Sensitive)
	case "auto_reply":
		config.CONFIG.AutoReply.Groups = groups
		viper.Set(key, config.CONFIG.AutoReply)
	case "online_course":
		config.CONFIG.OnlineCourse.Groups = groups
		viper.Set(key, config.CONFIG.OnlineCourse)
	default:
		return
	}
}
