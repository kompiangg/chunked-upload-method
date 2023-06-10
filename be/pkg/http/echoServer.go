package http

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/labstack/echo/v4"
)

type echoServer struct {
	port                 string
	environment          string
	whiteListAllowOrigin []string
	Echo                 *echo.Echo
}

type ServerConfig struct {
	Port                 string
	Environment          string
	WhiteListAllowOrigin []string
}

func InitEchoServer(config *ServerConfig) (echoServer, error) {
	if config == nil {
		return echoServer{}, errors.New("error while init echo server, config must be not nil")
	}

	return echoServer{
		port:                 config.Port,
		environment:          config.Environment,
		whiteListAllowOrigin: config.WhiteListAllowOrigin,
		Echo:                 echo.New(),
	}, nil
}

func (s *echoServer) ServeHTTP() error {
	port, err := strconv.Atoi(s.port)
	if err != nil {
		log.Println("[ERROR] Error while convert port to number:", err.Error())
		return err
	}

	errChan := make(chan error, 1)

	go func() {
		s.Echo.HideBanner = true
		if err := s.Echo.Start(fmt.Sprintf(":%d", port)); err != nil {
			errChan <- err
		}
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-errChan:
		log.Println("[ERROR] Error while starting server: ", err.Error())
		return err
	case <-signalChan:
		s.GracefulShutDown()
		return nil
	}
}

func (s *echoServer) GracefulShutDown() error {
	err := s.Echo.Server.Shutdown(context.Background())
	if err != nil {
		return err
	}

	return nil
}
