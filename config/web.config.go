package config

type WebConfig struct {
	Host string `json:"host" envconfig:"WEB_HOST"`
	Port string `json:"port" envconfig:"WEB_PORT"`
}
