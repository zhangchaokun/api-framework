package config

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

type app struct {
	UseDatabase bool
	UseRedis    bool
	UseMongo    bool
	UseRabbitMQ bool

	JwtSecret  string
	PrefixUrl  string
	PageSize   int
	TimeFormat string

	RuntimeRootPath string

	LogSavePath string
	LogSaveName string
	LogFileExt  string
}

type server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

type redis struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}


type mongo struct {
	URL      string
	Database string
}

type rabbit struct {
	RabbitURL string
}


var (
	AppConfig      = &app{}
	ServerConfig   = &server{}
	DatabaseConfig = &database{}
	RedisConfig    = &redis{}
	MongoConfig    = &mongo{}
	RabbitConfig   = &rabbit{}

	cfg *ini.File
)

func InitConfig() {
	var err error
	cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}

	iniMapTo("app", AppConfig)
	iniMapTo("server", ServerConfig)
	iniMapTo("database", DatabaseConfig)
	iniMapTo("redis", RedisConfig)
	iniMapTo("mongo", MongoConfig)
	iniMapTo("mongo", RabbitConfig)

	ServerConfig.ReadTimeout = ServerConfig.ReadTimeout * time.Second
	ServerConfig.WriteTimeout = ServerConfig.ReadTimeout * time.Second
	RedisConfig.IdleTimeout = RedisConfig.IdleTimeout * time.Second
}

func iniMapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("cfg.MapTo err: %v", err)
	}
}
