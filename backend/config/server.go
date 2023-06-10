package config

type ServerConfig struct {
	Port                 string   `validate:"required,notblank,numeric"`
	Environment          string   `validate:"required,notblank,oneof=dev prod"`
	WhiteListAllowOrigin []string `validate:"required,notblank"`
}
