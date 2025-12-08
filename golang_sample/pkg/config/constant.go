package config

// Config app
type Config struct {
	Port             string           `json:"port" mapstructure:"port"`
	ENV              string           `json:"env" mapstructure:"env"`
	Redis            RedisConfig      `json:"redis"  mapstructure:"redis"`
	Services         Service          `json:"services" mapstructure:"services"`
	LoginInformation LoginInformation `json:"login_information" mapstructure:"login_information"`
}

type RedisConfig struct {
	Address  string `json:"address" mapstructure:"address"`
	Password string `json:"password" mapstructure:"password"`
	DB       int    `json:"db" mapstructure:"db"`
}

type LoginInformation struct {
	ClientId     string `json:"client_id" mapstructure:"client_id"`
	ClientSecret string `json:"client_secret" mapstructure:"client_secret"`
}

// Service -
type Service struct {
	GPAYOpenAPI string `json:"gpay_open_api" mapstructure:"gpay_open_api"`
}
