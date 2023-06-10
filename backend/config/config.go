package config

import (
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
	"github.com/kompiangg/shipper-fp/pkg/validator"
)

type Config struct {
	DatabaseConfig   DatabaseConfig
	ServerConfig     ServerConfig
	SwaggerConfig    SwaggerConfig
	RedisConfig      RedisConfig
	UploadFileConfig UploadFileConfig
	CloudinaryConfig CloudinaryConfig
	UploadFolderPath string
}

func InitConfig(validator validator.ValidatorItf) (Config, error) {
	godotenv.Load("./etc/.env")

	config := Config{
		DatabaseConfig: DatabaseConfig{
			Username:     os.Getenv("DATABASE_USERNAME"),
			Password:     os.Getenv("DATABASE_PASSWORD"),
			Host:         os.Getenv("DATABASE_HOST"),
			Port:         os.Getenv("DATABASE_PORT"),
			DatabaseName: os.Getenv("DATABASE_NAME"),
		},
		ServerConfig: ServerConfig{
			Port:                 os.Getenv("SERVER_PORT"),
			Environment:          os.Getenv("SERVER_ENVIRONMENT"),
			WhiteListAllowOrigin: strings.Split(os.Getenv("SERVER_WHITE_LIST_ALLOW_ORIGIN"), ","),
		},
		SwaggerConfig: SwaggerConfig{
			Title:       os.Getenv("SWAGGER_TITLE"),
			Host:        os.Getenv("SWAGGER_HOST"),
			Port:        os.Getenv("SWAGGER_PORT"),
			Version:     os.Getenv("SWAGGER_VERSION"),
			Description: os.Getenv("SWAGGER_DESCRIPTION"),
			Schemes:     strings.Split(os.Getenv("SWAGGER_SCHEMES"), ","),
		},
		UploadFileConfig: UploadFileConfig{
			ChunkSizeChar: os.Getenv("FILE_CHUNK_SIZE"),
		},
		RedisConfig: RedisConfig{
			Hostname: os.Getenv("REDIS_HOSTNAME"),
			Port:     os.Getenv("REDIS_PORT"),
			Username: os.Getenv("REDIS_USERNAME"),
			Password: os.Getenv("REDIS_PASSWORD"),
			DBChar:   os.Getenv("REDIS_DB"),
		},
		CloudinaryConfig: CloudinaryConfig{
			APIKey:    os.Getenv("CLOUDINARY_API_KEY"),
			APISecret: os.Getenv("CLOUDINARY_API_SECRET"),
			CloudName: os.Getenv("CLOUDINARY_CLOUD_NAME"),
		},
	}

	err := validator.Validate(config)
	if err != nil {
		return config, err
	}

	config.UploadFileConfig.ChunkSize, err = strconv.Atoi(config.UploadFileConfig.ChunkSizeChar)
	if err != nil {
		return config, err
	}

	return config, nil
}
