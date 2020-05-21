package domain

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
