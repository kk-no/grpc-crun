package config

import (
	"github.com/kelseyhightower/envconfig"
)

var Conf *Config

type Config struct {
	Port string `envconfig:"PORT" default:"8000"`

	ServiceDomain string `envconfig:"SERVICE_DOMAIN" default:"localhost"`
	ServicePort string `envconfig:"SERVICE_PORT" default:"8080"`
}

func (c *Config) load() error {
	return envconfig.Process("", c)
}

func Load() error {
	conf := &Config{}
	if err := conf.load(); err != nil {
		return err
	}
	Conf = conf
	return nil
}