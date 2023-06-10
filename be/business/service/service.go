package service

import (
	"github.com/kompiangg/shipper-fp/business/repository"
	"github.com/kompiangg/shipper-fp/business/service/file"
	"github.com/kompiangg/shipper-fp/business/service/ping"
	"github.com/kompiangg/shipper-fp/config"
	"github.com/kompiangg/shipper-fp/pkg/validator"
)

type Service struct {
	Ping ping.PingItf
	File file.ServiceItf
}

func InitService(
	repository repository.Repository,
	config config.Config,
	validator validator.ValidatorItf,
) (Service, error) {
	pingService := ping.InitService(config)

	fileService, err := file.InitService(config, validator, repository.File, repository.Upload)
	if err != nil {
		return Service{}, err
	}

	return Service{
		Ping: &pingService,
		File: &fileService,
	}, nil
}
