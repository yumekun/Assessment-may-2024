package config

import (
	"fmt"

	"mutasi-service/utils/errs"

	"github.com/spf13/viper"
)

type Config struct {
	PostgresDriver string `mapstructure:"POSTGRES_DRIVER"`
	PostgresUrl    string `mapstructure:"POSTGRES_URL"`

	RedisServiceAddress string `mapstructure:"REDIS_SERVICE_ADDRESS"`
	RedisPassword       string `mapstructure:"REDIS_PASSWORD"`

	RedisMutasiRequestStream string `mapstructure:"REDIS_MUTASI_REQUEST_STREAM"`
}

func LoadConfig(path string) (config Config, err error) {
	const op errs.Op = "config/LoadConfig"

	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return config, errs.E(op, errs.IO, fmt.Sprintf("failed to read configuration file: %s", err))
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return config, errs.E(op, errs.Unanticipated, fmt.Sprintf("failed to unmarshal configuration: %s", err))
	}

	return
}
