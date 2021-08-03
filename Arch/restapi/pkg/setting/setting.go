package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

type App struct{}

var AppSetting = &App{}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

var DatabaseSetting = &Database{}

type Redis struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

var RedisSetting = &Redis{}

var cfg *ini.File

func Setup() {
	var err error
	cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("setting.Setup fail conf/app.ini %v", err)
	}

	if err = cfg.Section("app").MapTo(AppSetting); err != nil {
		log.Fatalf("Cfg.MapTo app err: %v", err)
	}
	if err = cfg.Section("server").MapTo(ServerSetting); err != nil {
		log.Fatalf("Cfg.MapTo server err: %v", err)
	}
	if err = cfg.Section("database").MapTo(DatabaseSetting); err != nil {
		log.Fatalf("Cfg.MapTo database err: %v", err)
	}
	if err = cfg.Section("redis").MapTo(RedisSetting); err != nil {
		log.Fatalf("Cfg.MapTo redis err: %v", err)
	}

	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second
	RedisSetting.IdleTimeout = RedisSetting.IdleTimeout * time.Second
}
