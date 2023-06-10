package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/kompiangg/shipper-fp/business/repository"
	"github.com/kompiangg/shipper-fp/business/service"
	"github.com/kompiangg/shipper-fp/cmd/webservice"
	"github.com/kompiangg/shipper-fp/config"
	"github.com/kompiangg/shipper-fp/pkg/db"
	"github.com/kompiangg/shipper-fp/pkg/db/sqlx"
	_ "github.com/kompiangg/shipper-fp/pkg/errors"
	"github.com/kompiangg/shipper-fp/pkg/objstorage/cloudinary"
	"github.com/kompiangg/shipper-fp/pkg/redis"
	"github.com/kompiangg/shipper-fp/pkg/validator"
)

//	@securityDefinitions.apikey	BearerAuth
//	@in							header
//	@name						Authorization
//	@description				Type "Bearer " before the token value
func main() {
	validator, err := validator.New()
	if err != nil {
		panic(err)
	}

	config, err := config.InitConfig(&validator)
	if err != nil {
		panic(err)
	}

	tmpDirPath, err := filepath.Abs("./tmp")
	if err != nil {
		panic(err)
	}

	if err := os.Mkdir(tmpDirPath, 0755); os.IsExist(err) {
		fmt.Println("The directory named", tmpDirPath, "exists")
	} else {
		fmt.Println("The directory named", tmpDirPath, "does not exist")
	}

	config.UploadFolderPath = tmpDirPath

	db, _, err := sqlx.Init(&sqlx.SqlxConfig{
		MySQLConfig: &db.MySQLConfig{
			Username:     config.DatabaseConfig.Username,
			Password:     config.DatabaseConfig.Password,
			Host:         config.DatabaseConfig.Host,
			Port:         config.DatabaseConfig.Port,
			DatabaseName: config.DatabaseConfig.DatabaseName,
			ParseTime:    true,
			Localization: "UTC",
		},
	})
	if err != nil {
		panic(err)
	}

	redis := redis.InitRedis(redis.RedisConfig{
		Hostname: fmt.Sprintf("%s:%s", config.RedisConfig.Hostname, config.RedisConfig.Port),
		Username: config.RedisConfig.Username,
		Password: config.RedisConfig.Password,
		DB:       config.RedisConfig.DB,
	})

	cloudinary, err := cloudinary.InitCloudinary(cloudinary.CloudinaryConfig{
		APIKey:    config.CloudinaryConfig.APIKey,
		APISecret: config.CloudinaryConfig.APISecret,
		CloudName: config.CloudinaryConfig.CloudName,
	})
	if err != nil {
		panic(err)
	}

	repository, err := repository.InitRepository(
		config,
		db,
		redis,
		cloudinary,
	)
	if err != nil {
		panic(err)
	}

	service, err := service.InitService(
		repository,
		config,
		&validator,
	)
	if err != nil {
		panic(err)
	}

	err = webservice.InitWebService(
		service,
		config,
	)
	if err != nil {
		panic(err)
	}
}
