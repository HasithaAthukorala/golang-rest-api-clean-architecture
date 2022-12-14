package main

import (
	"github.com/sirupsen/logrus"
	v1 "golang-rest-api-clean-architecture/pkg/api/v1"
	"golang-rest-api-clean-architecture/pkg/config"
	services "golang-rest-api-clean-architecture/pkg/external-services"
	"golang-rest-api-clean-architecture/pkg/server"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	stopCh := setupSignalHandler()
	cfg, err := config.Load("config")
	if err != nil {
		log.Fatalf("cannot read config file: %s", err.Error())
	}

	clientSet, _ := services.NewClients(cfg, stopCh)
	opt := server.Options{
		Port:       8080,
		EnableAuth: true,
	}
	routes := server.Routes{
		Prefix: "/api/v1",
		Routes: v1.Build(),
	}
	srv := server.New(clientSet, routes, opt)
	if err := srv.Run(stopCh); err != nil {
		logrus.Fatalf("Failed to run the api-server: %v", err)
	}
}

var onlyOneSignalHandler = make(chan struct{})
var shutdownSignals = []os.Signal{os.Interrupt, syscall.SIGTERM}

// SetupSignalHandler registered for SIGTERM and SIGINT. A stop channel is returned
// which is closed on one of these signals. If a second signal is caught, the program
// is terminated with exit code 1.
func setupSignalHandler() (stopCh <-chan struct{}) {
	close(onlyOneSignalHandler) // panics when called twice

	stop := make(chan struct{})
	c := make(chan os.Signal, 2)
	signal.Notify(c, shutdownSignals...)
	go func() {
		<-c
		close(stop)
		<-c
		os.Exit(1) // second signal. Exit directly.
	}()

	return stop
}
