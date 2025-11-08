package config

import (
	"flag"
	"gateway/internal/utils/format"
	"github.com/spf13/viper"
	"log"
	"time"
)

type Config struct {
	AddrProducts string        `mapstructure:"addr_products"`
	Timeout      time.Duration `mapstructure:"timeout"`
	Backoff      time.Duration `mapstructure:"backoff"`
	RetriesCount int           `mapstructure:"retries_count"`
	Port         int           `mapstructure:"port"`
	Mode         string        `mapstructure:"mode"`
	RedisPw      string        `mapstructure:"redis_pw"`
	RedisAddr    string        `mapstructure:"redis_addr"`
	AdminPW      string        `mapstructure:"admin_pw"`
}

// MustSetup return config and panic if error
func MustSetup() *Config {
	cfg, err := setup()
	if err != nil {
		log.Panic(err)
	}
	return cfg
}

// setup create config structure
func setup() (*Config, error) {
	const op = "config.setup"
	configPath := flag.String("config", "./config/local.yaml", "path to config file")
	flag.Parse()

	viper.SetConfigFile(*configPath)
	var cfg Config
	if err := viper.ReadInConfig(); err != nil {
		return nil, format.Error(op, err)
	}
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, format.Error(op, err)
	}
	if cfg.Mode == "DEV" {
		log.Println(format.Struct(cfg))
	}

	return &cfg, nil
}
