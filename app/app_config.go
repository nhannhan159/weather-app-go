package app

import (
	"io/ioutil"

	"github.com/kelseyhightower/envconfig"
	"github.com/nhannhan159/weather-app-go/domain"
	"gopkg.in/yaml.v2"
)

// https://dev.to/ilyakaznacheev/a-clean-way-to-pass-configs-in-a-go-application-1g64
func (app *App) initializeConfig() {
	var config domain.GlobalConfig

	readFile(&config)
	readEnv(&config)
	app.GlobalConfig = &config
}

func readFile(config *domain.GlobalConfig) {
	f, err := ioutil.ReadFile("/Users/tien.tan/go/src/github.com/nhannhan159/weather-app-go/config/config.yml")
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(f, config)
	if err != nil {
		panic(err)
	}
}

func readEnv(config *domain.GlobalConfig) {
	err := envconfig.Process("", config)
	if err != nil {
		panic(err)
	}
}
