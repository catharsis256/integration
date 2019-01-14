package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

/**
https://medium.com/seek-blog/microservices-in-go-2fc1570f6800
https://github.com/go-kit/kit/tree/master/examples/addsvc/pkg
http://www.ru-rocker.com/2017/04/17/micro-services-using-go-kit-service-discovery/
https://github.com/peterbourgon/go-microservices

*/

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
		// ... a(ctx)
		println("SupervisorHolder initiated")
	}()
}
