package config

type DBConfig struct {
	Host     string `json:"host" envconfig:"DB_HOST"`
	Port     string `json:"port" envconfig:"DB_PORT"`
	Database string `json:"database" envconfig:"DB_DATABASE"`
	Username string `json:"username" envconfig:"DB_USERNAME"`
	Password string `json:"password" envconfig:"DB_PASSWORD"`
}
