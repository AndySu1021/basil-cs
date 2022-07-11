package config

import (
	"github.com/AndySu1021/go-util/db"
	"github.com/AndySu1021/go-util/gin"
	"github.com/AndySu1021/go-util/logger"
	"github.com/AndySu1021/go-util/redis"
	"github.com/AndySu1021/go-util/storage"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

type Config struct {
	Salt              string   `mapstructure:"salt"`
	AdminPassword     string   `mapstructure:"admin_password"`
	WebChatBaseUrl    string   `mapstructure:"web_chat_base_url"`
	WsOriginWhiteList []string `mapstructure:"ws_origin_white_list"`
}

// AppConfig APP設定
type AppConfig struct {
	fx.Out

	App      *Config           `mapstructure:"app"`
	Storage  *storage.Config   `mapstructure:"storage"`
	Http     *gin.Config       `mapstructure:"http"`
	Log      *logger.ZapConfig `mapstructure:"log"`
	Database *db.Config        `mapstructure:"database"`
	Mongo    *db.MongoConfig   `mapstructure:"mongo"`
	Redis    *redis.Config     `mapstructure:"redis"`
}

// NewConfig Initiate config
func NewConfig() (AppConfig, error) {
	viper.AutomaticEnv()
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	var config = AppConfig{}

	if err := viper.ReadInConfig(); err != nil {
		return AppConfig{}, err
	}
	if err := viper.Unmarshal(&config); err != nil {
		return AppConfig{}, err
	}

	return config, nil
}
