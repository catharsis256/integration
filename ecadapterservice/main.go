package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

/**
Following articles should help me develop this project:

https://medium.com/seek-blog/microservices-in-go-2fc1570f6800
https://github.com/go-kit/kit/tree/master/examples/addsvc/pkg
https://github.com/peterbourgon/go-microservices
http://callistaenterprise.se/blogg/teknik/2017/02/21/go-blog-series-part2/
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
