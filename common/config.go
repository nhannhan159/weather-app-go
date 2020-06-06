package common

type ServerConfig struct {
	Port string `yaml:"port" envconfig:"PORT"`
	Host string `yaml:"host" envconfig:"HOST"`
}

type GRPCConfig struct {
	Port    string      `yaml:"port" envconfig:"GRPC_PORT"`
	Clients GRPCClients `yaml:"clients"`
}

type GRPCClients struct {
	CityService    string `yaml:"city_service" envconfig:"GRPC_CITY_SERVICE"`
	WeatherService string `yaml:"weather_service" envconfig:"GRPC_WEATHER_SERVICE"`
}

type DatabaseConfig struct {
	Dialect string `yaml:"dialect" envconfig:"DIALECT"`
	DSN     string `yaml:"dsn" envconfig:"DB_DSN"`
}

type GlobalConfig struct {
	AppName  string         `yaml:"app-name" envconfig:"APP_NAME"`
	BaseDir  string         `yaml:"base-dir" envconfig:"BASE_DIR"`
	GRPC     GRPCConfig     `yaml:"grpc"`
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
}
