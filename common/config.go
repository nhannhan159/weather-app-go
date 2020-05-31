package common

type ServerConfig struct {
	Port string `yaml:"port" envconfig:"SERVER_PORT"`
	Host string `yaml:"host" envconfig:"SERVER_HOST"`
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
	GRPC     GRPCConfig     `yaml:"grpc"`
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
}
