package config

import (
	"gocarch/pkg/constants"
	"time"
)

type (
	MainConfig struct {
		Service     ServiceConfig `yaml:"Service"`
		Server      ServerConfig  `yaml:"Server"`
		InternalAPI APIConfig     `yaml:"InternalAPI"`
	}

	ServiceConfig struct {
		Name string `yaml:"Name"`
	}

	ServerConfig struct {
		Port            string        `yaml:"Port"`
		GracefulTimeout time.Duration `yaml:"GracefulTimeout"`
		ReadTimeout     time.Duration `yaml:"ReadTimeout"`
		WriteTimeout    time.Duration `yaml:"WriteTimeout"`
	}

	APIConfig struct {
		BasePath      string `yaml:"BasePath"`
		APITimeout    int    `yaml:"APITimeout"`
		EnableSwagger bool   `yaml:"EnableSwagger" default:"true"`
	}
)

func ReadModuleConfig(cfg interface{}, module, configLocation string) interface{} {
	if configLocation == constants.EmptyString {
		configLocation = "config/files"
	}

	// if err := configreader.ReadModuleConfig(cfg, configLocation, module); err != nil {
	// 	log.Fatalln("failed to read config for ", module)
	// }

	return cfg
}
