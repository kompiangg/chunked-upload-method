package config

type SwaggerConfig struct {
	Title       string   `validate:"required,notblank"`
	Host        string   `validate:"required,notblank"`
	Port        string   `validate:"required,notblank,numeric"`
	Version     string   `validate:"required,notblank"`
	Description string   `validate:"required,notblank"`
	Schemes     []string `validate:"required,notblank,dive,notblank"`
}
