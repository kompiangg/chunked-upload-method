package ping

import (
	"os"
	"time"
)

func (p *Ping) Ping() string {
	return "pong from " + p.config.ServerConfig.Environment + " at " + time.Now().String() + " ( " + os.Getenv("POD_NAME") + " )"
}
