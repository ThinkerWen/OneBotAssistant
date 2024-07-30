package config

import (
	"errors"
	"github.com/charmbracelet/log"
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	Proxy        string             `mapstructure:"proxy"`
	Hosts        []int64            `mapstructure:"hosts"`
	ApiUrl       string             `mapstructure:"api_url"`
	Molly        MollyConfig        `mapstructure:"molly"`
	HeroPower    HeroPowerConfig    `mapstructure:"hero_power"`
	OnlineCourse OnlineCourseConfig `mapstructure:"online_course"`
}

type HeroPowerConfig struct {
	Enable bool    `mapstructure:"enable"`
	Groups []int64 `mapstructure:"groups"`
}

type MollyConfig struct {
	Enable    bool    `mapstructure:"enable"`
	QQ        int64   `mapstructure:"qq"`
	Name      string  `mapstructure:"name"`
	ApiKey    string  `mapstructure:"api_key"`
	ApiSecret string  `mapstructure:"api_secret"`
	Groups    []int64 `mapstructure:"groups"`
}

type OnlineCourseConfig struct {
	Enable bool    `mapstructure:"enable"`
	Token  string  `mapstructure:"token"`
	Limit  int     `mapstructure:"limit"`
	Groups []int64 `mapstructure:"groups"`
}

var CONFIG Config

func init() {
	workDir, _ := os.Getwd()
	viper.AddConfigPath(workDir)
	viper.SetConfigType("yaml")
	viper.SetConfigName("application")
	initDefaultConfig()

	if err := viper.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			if err = viper.SafeWriteConfig(); err != nil {
				return
			}
		}
	}

	if err := viper.Unmarshal(&CONFIG); err != nil {
		return
	}

	if err := saveConfig(); err != nil {
		log.Error("Error saving config: ", err)
	}
}

func initDefaultConfig() {
	viper.SetDefault("proxy", "")
	viper.SetDefault("hosts", []int64{})
	viper.SetDefault("api_url", "ws://127.0.0.1:3001")

	viper.SetDefault("hero_power", HeroPowerConfig{
		Enable: true,
		Groups: []int64{},
	})
	viper.SetDefault("molly", MollyConfig{
		Enable:    true,
		QQ:        123456,
		Name:      "molly-bot",
		ApiKey:    "",
		ApiSecret: "",
		Groups:    []int64{},
	})
	viper.SetDefault("online_course", OnlineCourseConfig{
		Enable: true,
		Token:  "free",
		Limit:  1,
		Groups: []int64{},
	})
}

// saveConfig 将配置保存回文件
func saveConfig() error {
	if _, err := os.Stat(viper.ConfigFileUsed()); os.IsNotExist(err) {
		if err = viper.WriteConfigAs("application.yaml"); err != nil {
			return err
		}
	} else {
		if err = viper.WriteConfig(); err != nil {
			return err
		}
	}

	log.Info("Configuration saved successfully.")
	return nil
}
