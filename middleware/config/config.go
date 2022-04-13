package config

import (
	"codebase/go-codebase/config"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Cfg config.Config
}

func CreateConfig(logger *logrus.Logger) *Config {
	return &Config{
		Cfg: config.CreateGlobalConfig(logger),
	}
}

func (c *Config) Redis(address, password string) *redis.Client {
	return c.Cfg.Redis(address, password)
}
