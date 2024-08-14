package util

import (
	"OneBotAssistant/config"
	"github.com/charmbracelet/log"
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

func AddHost(host int64) {
	if IsHost(host) {
		return
	}
	config.CONFIG.Hosts = append(config.CONFIG.Hosts, host)
	if err := config.SaveConfig(); err != nil {
		log.Error("Save config fail", "err", err)
	}
}

func DelHost(host int64) {
	var result []int64
	for _, v := range config.CONFIG.Hosts {
		if v != host {
			result = append(result, v)
		}
	}
	config.CONFIG.Hosts = result
	if err := config.SaveConfig(); err != nil {
		log.Error("Error writing config", "err", err)
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
		log.Error("Error writing config", "err", err)
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
		log.Error("Error writing config", "err", err)
	}
}

func setConfig(groups []int64, key string) {
	switch key {
	case "molly":
		config.CONFIG.Molly.Groups = groups
	case "hero_power":
		config.CONFIG.HeroPower.Groups = groups
	case "sensitive":
		config.CONFIG.Sensitive.Groups = groups
	case "auto_reply":
		config.CONFIG.AutoReply.Groups = groups
	case "online_course":
		config.CONFIG.OnlineCourse.Groups = groups
	default:
		return
	}
}
