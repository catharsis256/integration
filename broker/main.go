package broker

import (
	"context"
	"gitlab.smartblocklab.com/integration/broker/service"
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
		service.NewSupervisorHolder(service.NewSupervisorService()).GetService().Initiate(ctx)
		println("SupervisorHolder initiated")
	}()
}
