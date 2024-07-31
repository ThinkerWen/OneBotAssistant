package config

import (
	"database/sql"
	"errors"
	"github.com/charmbracelet/log"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	Proxy        string             `mapstructure:"proxy"`
	Hosts        []int64            `mapstructure:"hosts"`
	ApiUrl       string             `mapstructure:"api_url"`
	Molly        MollyConfig        `mapstructure:"molly"`
	HeroPower    HeroPowerConfig    `mapstructure:"hero_power"`
	Sensitive    SensitiveConfig    `mapstructure:"sensitive"`
	AutoReply    AutoReplyConfig    `mapstructure:"auto_reply"`
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
	Groups    []int64 `mapstructure:"groups"`
	ApiKey    string  `mapstructure:"api_key" json:"api_key"`
	ApiSecret string  `mapstructure:"api_secret" json:"api_secret"`
}

type OnlineCourseConfig struct {
	Enable bool    `mapstructure:"enable"`
	Token  string  `mapstructure:"token"`
	Limit  int     `mapstructure:"limit"`
	Groups []int64 `mapstructure:"groups"`
}

type SensitiveConfig struct {
	Enable      bool    `mapstructure:"enable"`
	Groups      []int64 `mapstructure:"groups"`
	AlertTimes  int     `mapstructure:"alert_times" json:"alert_times"`
	ShutSeconds int     `mapstructure:"shut_seconds" json:"shut_seconds"`
}

type AutoReplyConfig struct {
	Enable bool    `mapstructure:"enable"`
	Groups []int64 `mapstructure:"groups"`
}

var DB *sql.DB
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

	if err := checkAndCreateDatabase(); err != nil {
		log.Error(err)
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
	viper.SetDefault("sensitive", SensitiveConfig{
		Enable:      true,
		Groups:      []int64{},
		AlertTimes:  3,
		ShutSeconds: 60,
	})
	viper.SetDefault("auto_reply", AutoReplyConfig{
		Enable: true,
		Groups: []int64{},
	})
}

// checkAndCreateDatabase 初始化SQLite表
func checkAndCreateDatabase() error {
	var err error
	if DB, err = sql.Open("sqlite3", "one_bot_assistant.db"); err != nil {
		return err
	}

	createAutoReply := `
	CREATE TABLE IF NOT EXISTS "auto_reply" (
		id         integer 		not null primary key autoincrement,
		ask        varchar(255) default ''                not null,
		reply      TEXT         default ''                not null,
		group_id   varchar(15)  default ''                not null,
		user_id    varchar(15)  default ''                not null,
		created_at timestamp    default CURRENT_TIMESTAMP not null
	);`
	createSensitive := `
	CREATE TABLE IF NOT EXISTS "sensitive_words" (
	  "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	  "sensitive_word" varchar(255) NOT NULL DEFAULT '',
	  "group_id" varchar(15) NOT NULL DEFAULT '',
	  "user_id" varchar(15) NOT NULL DEFAULT '',
	  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
	);`

	if _, err = DB.Exec(createAutoReply); err != nil {
		return err
	}
	if _, err = DB.Exec(createSensitive); err != nil {
		return err
	}

	return nil
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
