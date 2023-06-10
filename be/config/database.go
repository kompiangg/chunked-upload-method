package config

type DatabaseConfig struct {
	Username     string `validate:"required,notblank"`
	Password     string `validate:"required,notblank"`
	Host         string `validate:"required,notblank"`
	Port         string `validate:"required,notblank,numeric"`
	DatabaseName string `validate:"required,notblank"`
}
