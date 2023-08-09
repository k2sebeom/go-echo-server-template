package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/k2sebeom/go-echo-server-template/config"
	"github.com/k2sebeom/go-echo-server-template/server"
)

func main() {
	config.InitializeConfig("./envs")

	s, err := server.NewServer()
	if err != nil {
		log.Fatal(err.Error())
	}

	go func() {
		log.Printf("Starting Care25 server on port %s", s.Config.Port)
		if err := s.Echo.Start(fmt.Sprintf(":%s", s.Config.Port)); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
