package config

type RedisConfig struct {
	Hostname string
	Username string
	Port     string
	Password string
	DBChar   string

	DB int
}
