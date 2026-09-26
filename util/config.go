package util

import (
	"errors"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DBDriver            string        `mapstructure:"DB_DRIVER"`
	DBSource            string        `mapstructure:"DB_SOURCE"`
	ServerAddress       string        `mapstructure:"SERVER_ADDRESS"`
	TokenSymmetricKey   string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)

	viper.SetDefault("DB_DRIVER", "postgres")
	viper.SetDefault("DB_SOURCE",
		"postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable",
	)

	for _, key := range []string{
		"DB_DRIVER",
		"DB_SOURCE",
		"SERVER_ADDRESS",
		"TOKEN_SYMMETRIC_KEY",
		"ACCESS_TOKEN_DURATION",
	} {
		if err := viper.BindEnv(key); err != nil {
			return config, err
		}
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		var configFileNotFound viper.ConfigFileNotFoundError
		if !errors.As(err, &configFileNotFound) {
			return config, err
		}
		// A local app.env file is optional; environment variables/defaults are sufficient.
	}

	err = viper.Unmarshal(&config)
	return config, err
}
