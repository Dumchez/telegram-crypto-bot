package config

import "github.com/spf13/viper"

type Config struct {
	TelegramToken string
}

func InitConfig() (*Config, error) {
	var cfg Config
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	cfg.TelegramToken = viper.GetString("token")
	return &cfg, nil
}
