package config

import (
	"errors"
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"time"
)

type Config struct {
	Env        string     `yaml:"env" env-required:"true"`
	HttpServer HttpServer `yaml:"http_server"`
	Database   Database   `yaml:"database"`
}
type HttpServer struct {
	Address     string        `yaml:"address" env-required:"true"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}
type Database struct {
	Host     string `yaml:"host" env-default:"localhost" env-required:"true"`
	Port     string `yaml:"port" env-default:"" env-required:"true"`
	Username string `yaml:"username" env-default:"marat" env-required:"true"`
	Password string `yaml:"password" emv-required:"true"`
	DBName   string `yaml:"db_name" env-default:"Test" env-required:"true"`
}

func Load(confPath string) (*Config, error) {
	if confPath == "" {
		return nil, errors.New("The path argument for the configuration file is not specified!")

	}
	_, err := os.Stat(confPath)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Config file %s does not exist!", confPath))

	}

	var cfg Config
	err = cleanenv.ReadConfig(confPath, &cfg)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Cannot read config %s", err))

	}
	return &cfg, nil
}
