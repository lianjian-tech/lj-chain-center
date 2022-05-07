package pkg

import (
	"github.com/fsnotify/fsnotify"
	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"lj-chain-center/common/log"
)

var (
	Cfg      *Config
	DB       *sqlx.DB
	RedisCli *redis.ClusterClient
)

const (
	YYYY_MM_DD          = "2006-01-02"
	YYYY_MM_DD_HH_MM_SS = "2006-01-02 15:04:05"
	CORE_TYPE           = 0
	ANT_BASS_CHAIN_TYPE = 1
	LUBAN_CHAIN_TYPE    = 2
	COMPANY_TYPE        = 1
	PERSON_TYPE         = 2
	VALID               = 1
	INVALID             = 0
	HANDLE_ERROR        = "handle error"
)

type Config struct {
	Logger  *LoggerConfig
	AntBass *AntBassConfig
	Luban   *LubanConfig
}

func Init(cfgName string) {
	setConfig(cfgName)
	Cfg = loadConfig()
	initConfig(Cfg)
	watchConfig()
}

func setConfig(cfgName string) {
	if cfgName != "" {
		viper.SetConfigFile(cfgName)
	} else {
		viper.AddConfigPath("conf")
		viper.SetConfigName("config-local")
	}
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		panic("initConfig error")
	}
}

func loadConfig() *Config {
	cfg := &Config{
		Logger:  LoadLoggerConfig(viper.Sub("logger")),
		Luban:   LoadLubanConfig(viper.Sub("luban")),
		AntBass: LoadAntBassConfig(viper.Sub("ant_bass"))}
	return cfg
}

func initConfig(cfg *Config) {
	cfg.Logger.InitLogger()
}

func watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Infof("Config file changed: %s", e.Name)
	})
}

type AntBassConfig struct {
	Url string
}

func LoadAntBassConfig(viper *viper.Viper) *AntBassConfig {
	cfg := &AntBassConfig{
		Url: viper.GetString("url"),
	}
	return cfg
}

type LubanConfig struct {
	Url string
}

func LoadLubanConfig(viper *viper.Viper) *LubanConfig {
	cfg := &LubanConfig{
		Url: viper.GetString("url"),
	}
	return cfg
}

type AuthConfig struct {
	PublicKey string
	Address   string
}

type BifConfig struct {
	BaseUrl   string
	ApiKey    string
	ApiSecret string
}
