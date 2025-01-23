package main

import __sealights__ "github.com/liornabat-sealights/go-calc-demo/__sealights__"

import (
	"context"
	"github.com/liornabat-sealights/go-calc-demo/service/server"
	"os"
	"os/signal"
	"syscall"
)

func setup(shutdown chan struct{}) {
	__sealights__.TraceFunc("d1334dab5a7a38cb6a")
	s := server.NewServer()
	if err := s.Start(context.Background()); err != nil {
		panic(err)
	}
	<-shutdown

}
func main() {
	__sealights__.StartMainFunc("7084e712c6351d4e37")
	defer func() { __sealights__.EndMainFunc("7084e712c6351d4e37") }()
	shutdown := make(chan struct{})
	go setup(shutdown)
	var gracefulShutdown = make(chan os.Signal, 1)
	signal.Notify(gracefulShutdown, syscall.SIGTERM)
	signal.Notify(gracefulShutdown, syscall.SIGINT)
	signal.Notify(gracefulShutdown, syscall.SIGQUIT) // wait for SIGKill
	<-gracefulShutdown
	close(shutdown)
}
