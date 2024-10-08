package config

import (
	"strings"
	"sync"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type (
	Config struct {
		Server  *Server  `mapstructure:"server" validate:"required"`
		MariaDB *MariaDB `mapstructure:"mariadb" validate:"required"`
		MongoDB *MongoDB `mapstructure:"mongodb" validate:"required"`
	}

	Server struct {
		Port         int      `mapstructure:"port" validate:"required"`
		AllowOrigins []string `mapstructure:"allowOrigins" validate:"required"`
		BodyLimit    string   `mapstructure:"bodyLimit" validate:"required"`
		Timeout      int      `mapstructure:"timeout" validate:"required"`
	}

	MariaDB struct {
		User     string `mapstructure:"user" validate:"required"`
		Password string `mapstructure:"password" validate:"required"`
		Database string `mapstructure:"database" validate:"required"`
		Host     string `mapstructure:"host" validate:"required"`
		Port     int    `mapstructure:"port" validate:"required"`
	}

	MongoDB struct {
		User     string `mapstructure:"user" validate:"required"`
		Password string `mapstructure:"password" validate:"required"`
		Host     string `mapstructure:"host" validate:"required"`
		Port     int    `mapstructure:"port" validate:"required"`
	}
)

var (
	once   sync.Once
	config *Config
)

func ConfigGetting() *Config {
	once.Do(func() {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("./config")
		viper.AutomaticEnv()
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
		if err := viper.ReadInConfig(); err != nil {
			panic(err)
		}

		if err := viper.Unmarshal(&config); err != nil {
			panic(err)
		}

		validating := validator.New()
		if err := validating.Struct(config); err != nil {
			panic(err)
		}
	})
	return config
}
