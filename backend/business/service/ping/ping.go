package ping

import "github.com/kompiangg/shipper-fp/config"

type PingItf interface {
	Ping() string
}

type Ping struct {
	config config.Config
}

func InitService(
	config config.Config,
) Ping {
	return Ping{
		config: config,
	}
}
