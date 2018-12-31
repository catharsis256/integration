package main

import (
	"com.smartblocklab/integration/core/service"
	"context"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	c := make(chan os.Signal, 2)

	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	defer func() {
		signal.Stop(c)
		cancel()
	}()

	go func() {
		select {
		case <-c:
			cancel()
			println("Context cancel completed")
		case <-ctx.Done():
			println("Program finished")
		}
	}()

	go func() {
		proxy := service.NewSupervisorProxy(service.NewSupervisorService())
		if sInterface, ok := proxy.(*service.SupervisorProxy); ok {
			sInterface.Initiate(ctx)
			println("SupervisorProxy initiated")
		}
	}()
}
