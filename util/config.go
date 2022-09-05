package util

import "github.com/spf13/viper"

type Config struct {
	DBUserName          string `mapstructure:"DB_USERNAME"`
	DBPassWord          string `mapstructure:"DB_PASSWORD"`
	DBHost              string `mapstructure:"DB_HOST"`
	DBName              string `mapstructure:"DB_NAME"`
	URL_WEBHOOK_DISCORD string `mapstructure:"URL_WEBHOOK_DISCORD"`
	TYPESENSE_URL       string `mapstructure:"TYPESENSE_URL"`
	TYPESENSE_API_KEY   string `mapstructure:"TYPESENSE_API_KEY"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
