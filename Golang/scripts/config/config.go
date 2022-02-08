package config

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

////////////////////
// 1、有一个注册了就不能继续注册了?
////////////////////

type (
	Config struct {
		AppName  string `mapstructure:"app_name"`
		LogLevel string `mapstructure:"log_level"`

		Server ServerConfig `mapstructure:"server"`
		DATA   DBConfig     `mapstructure:"data"`
		MySQL  MySQLConfig  `mapstructure:"mysql"`
		Redis  RedisConfig  `mapstructure:"redis"`
	}

	ServerConfig struct {
		BuildInfo    string `mapstructure:"build_info"`
		Addr         string `mapstructure:"addr"`
		SessionPath  string `mapstructure:"session_path"`
		ResetAdmin   string `mapstructure:"reset_admin"`
		Oauth2Server string `mapstructure:"oauth2_server"`
		Version      bool   `mapstructure:"version"`
	}

	DBConfig struct {
		Driver string `mapstructure:"driver"`
		DSN    string `mapstructure:"DSN"`
	}

	MySQLConfig struct {
		IP       string `mapstructure:"ip"`
		Port     int    `mapstructure:"port"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Database string `mapstructure:"database"`
	}

	RedisConfig struct {
		IP   string `mapstructure:"ip"`
		Port int    `mapstructure:"port"`
	}
)

var (
	Conf = new(Config)

	cliConf     = viper.New()
	defaultConf = viper.New()
	prodConf    = viper.New()
)

// cli
func init() {
	buildInfo := ""
	pflag.String("product", "", "应用模式")
	pflag.String("s", "", "server address")
	pflag.String("sp", "", "session path")
	pflag.String("db", "", "database path")
	pflag.String("resetAdmin", "", "reset admin new password")
	pflag.String("o2", "", "oauth2 server")
	pflag.Bool("v", false, buildInfo)

	cliConf.RegisterAlias("server.addr", "s")
	cliConf.RegisterAlias("server.session_path", "sp")
	cliConf.RegisterAlias("data.dsn", "db")
	cliConf.RegisterAlias("server.reset_admin", "resetAdmin")
	cliConf.RegisterAlias("server.oauth2_server", "o2")

	cliConf.BindPFlags(pflag.CommandLine)
	pflag.Parse()
}

// product
func init() {
	// 读取环境变量中dev or prod
	var file string
	if os.Getenv("material") != "" {
		file = fmt.Sprintf("config.%s.yaml", os.Getenv("material"))
	} else if cliConf.Get("product") != "" {
		file = fmt.Sprintf("config.%s.yaml", cliConf.Get("product"))
	}
	_, err := os.Stat(file)
	if os.IsNotExist(err) {
		return
	}
	prodConf.SetConfigFile(file)
	prodConf.AddConfigPath(".")
	err = prodConf.ReadInConfig()
	if err != nil {
		log.Fatalf("read config failed: %v", err)
	}
}

// default
func init() {
	// 读取默认配置
	_, err := os.Stat("config.yaml")
	if os.IsNotExist(err) {
		return
	}
	defaultConf.SetConfigFile("config.yaml")
	defaultConf.AddConfigPath(".")
	err = defaultConf.ReadInConfig()
	if err != nil {
		log.Fatalf("read config failed: %v", err)
	}
}

// ConfigLoad 装载 好像不能合并
func init() {
	confs := []*viper.Viper{defaultConf, prodConf, cliConf}
	for _, conf := range confs {
		for _, key := range conf.AllKeys() {
			if defaultConf.IsSet(key) && conf.IsSet(key) {
				viper.Set(key, conf.Get(key))
			}
		}
	}
	viper.Unmarshal(&Conf)
}
