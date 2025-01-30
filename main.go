package main

import (
	"context"
	"github.com/liornabat-sealights/go-calc-demo/service/server"
	"os"
	"os/signal"
	"syscall"
)

func setup(shutdown chan struct{}) {

	s := server.NewServer()
	if err := s.Start(context.Background()); err != nil {
		panic(err)
	}
	<-shutdown

}
func main() {
	shutdown := make(chan struct{})
	go setup(shutdown)
	var gracefulShutdown = make(chan os.Signal, 1)
	signal.Notify(gracefulShutdown, syscall.SIGTERM)
	signal.Notify(gracefulShutdown, syscall.SIGINT)
	signal.Notify(gracefulShutdown, syscall.SIGQUIT) // wait for SIGKill
	<-gracefulShutdown
	close(shutdown)
}
