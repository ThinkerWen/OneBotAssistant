package config

import (
	"database/sql"
	"github.com/charmbracelet/log"
	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/yaml.v3"
	"os"
	"sync"
)

type Config struct {
	ApiUrl       string             `yaml:"api_url"`
	Proxy        string             `yaml:"proxy"`
	Hosts        []int64            `yaml:"hosts"`
	Molly        MollyConfig        `yaml:"molly"`
	HeroPower    HeroPowerConfig    `yaml:"hero_power"`
	Sensitive    SensitiveConfig    `yaml:"sensitive"`
	AutoReply    AutoReplyConfig    `yaml:"auto_reply"`
	OnlineCourse OnlineCourseConfig `yaml:"online_course"`
}

type HeroPowerConfig struct {
	Enable bool    `yaml:"enable"`
	Groups []int64 `yaml:"groups"`
}

type MollyConfig struct {
	Enable    bool    `yaml:"enable"`
	QQ        int64   `yaml:"qq"`
	Name      string  `yaml:"name"`
	Groups    []int64 `yaml:"groups"`
	ApiKey    string  `yaml:"api_key"`
	ApiSecret string  `yaml:"api_secret"`
}

type OnlineCourseConfig struct {
	Enable bool    `yaml:"enable"`
	Token  string  `yaml:"token"`
	Limit  int     `yaml:"limit"`
	Groups []int64 `yaml:"groups"`
}

type SensitiveConfig struct {
	Enable      bool    `yaml:"enable"`
	Groups      []int64 `yaml:"groups"`
	AlertTimes  int     `yaml:"alert_times"`
	ShutSeconds int     `yaml:"shut_seconds"`
}

type AutoReplyConfig struct {
	Enable bool    `yaml:"enable"`
	Groups []int64 `yaml:"groups"`
}

var DB *sql.DB
var CONFIG Config
var mu sync.Mutex

func init() {
	workDir, _ := os.Getwd()
	yamlFile, err := os.ReadFile(workDir + "/application.yaml")
	if err != nil {
		initDefaultConfig()
		if err = SaveConfig(); err != nil {
			log.Fatalf("Error saving config: %v", err)
		}
	} else if err = yaml.Unmarshal(yamlFile, &CONFIG); err != nil {
		log.Fatalf("Error unmarshalling YAML: %v", err)
	}

	if err = checkAndCreateDatabase(); err != nil {
		log.Fatalf("Error creating database: %v", err)
	}
	log.Info("Load config successfully")
}

func initDefaultConfig() {
	config := new(Config)
	config.Proxy = ""
	config.Hosts = []int64{}
	config.ApiUrl = "ws://127.0.0.1:3001"

	config.HeroPower = HeroPowerConfig{
		Enable: true,
		Groups: []int64{},
	}

	config.Molly = MollyConfig{
		Enable:    true,
		QQ:        123456,
		Name:      "molly-bot",
		ApiKey:    "",
		ApiSecret: "",
		Groups:    []int64{},
	}

	config.OnlineCourse = OnlineCourseConfig{
		Enable: true,
		Token:  "free",
		Limit:  1,
		Groups: []int64{},
	}

	config.Sensitive = SensitiveConfig{
		Enable:      true,
		Groups:      []int64{},
		AlertTimes:  3,
		ShutSeconds: 60,
	}

	config.AutoReply = AutoReplyConfig{
		Enable: true,
		Groups: []int64{},
	}

	CONFIG = *config
	log.Info("Set default config successfully")
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

// SaveConfig 将配置保存回文件
func SaveConfig() error {
	mu.Lock()
	defer mu.Unlock()
	workDir, _ := os.Getwd()
	yamlData, err := yaml.Marshal(CONFIG)
	if err != nil {
		return err
	}

	if err = os.WriteFile(workDir+"/application.yaml", yamlData, 0644); err != nil {
		return err
	}

	log.Info("Configuration saved successfully")
	return nil
}
