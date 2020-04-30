package setting

import (
	"time"

	"github.com/go-ini/ini"
	log "github.com/sirupsen/logrus"
)

//-----------------------------------------------------------------------------

// App setting
type App struct {
	JwtSecret   string
	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string
}

// AppSetting app setting
var AppSetting = &App{}

// Server setting
type Server struct {
	RunMode      string
	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

// ServerSetting server setting
var ServerSetting = &Server{}

// Database setting
type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

// DatabaseSetting database setting
var DatabaseSetting = &Database{}

// Redis setting
type Redis struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

// RedisSetting redis setting
var RedisSetting = &Redis{}

var cfg *ini.File

// Setup 配置实例化
func Setup() {
	var err error
	cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'conf/app.ini': %v", err)
	}

	mapTo("app", AppSetting)
	mapTo("server", ServerSetting)
	mapTo("database", DatabaseSetting)
	mapTo("redis", RedisSetting)

	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second
	RedisSetting.IdleTimeout = RedisSetting.IdleTimeout * time.Second
}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s error: %v", section, err)
	}
}
