package config

import (
	"fmt"
	"os"

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
	f, err := os.Open("config.yml")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
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
