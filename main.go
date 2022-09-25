package main

import (
	"go.uber.org/zap"
	v1 "golang-rest-api-clean-architecture/pkg/api/v1"
	services "golang-rest-api-clean-architecture/pkg/external-services"
	"golang-rest-api-clean-architecture/pkg/server"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	stopCh := setupSignalHandler()
	clientSet, _ := services.NewClients()
	opt := server.Options{
		Port:       8080,
		EnableAuth: true,
	}
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugaredLogger := logger.Sugar()
	routes := server.Routes{
		Prefix: "/api/v1",
		Routes: v1.Build(),
	}
	srv := server.New(clientSet, routes, opt, sugaredLogger)
	if err := srv.Run(stopCh); err != nil {
		sugaredLogger.Fatalf("Failed to run the api-server: %v", err)
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
