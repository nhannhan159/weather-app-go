package config

import (
	"fmt"
	"io/ioutil"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
)

type ServerConfig struct {
	Port string `yaml:"port" envconfig:"SERVER_PORT"`
	Host string `yaml:"host" envconfig:"SERVER_HOST"`
}

type DatabaseConfig struct {
	Dialect string `yaml:"dialect" envconfig:"DIALECT"`
	DSN     string `yaml:"dsn" envconfig:"DB_DSN"`
}

type GlobalConfig struct {
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
}

// https://dev.to/ilyakaznacheev/a-clean-way-to-pass-configs-in-a-go-application-1g64
func InitializeConfig() GlobalConfig {
	var cfg GlobalConfig
	readFile(&cfg)
	readEnv(&cfg)
	fmt.Printf("%+v", cfg)
	return cfg
}

func readFile(cfg *GlobalConfig) {
	f, err := ioutil.ReadFile("/Users/tien.tan/go/src/github.com/nhannhan159/weather-app-go/config/config.yml")
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(f, cfg)
	if err != nil {
		panic(err)
	}
}

func readEnv(cfg *GlobalConfig) {
	err := envconfig.Process("", cfg)
	if err != nil {
		panic(err)
	}
}
