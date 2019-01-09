package server

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

type GrpcServerBuilder struct {
	server   *grpc.Server
	listener net.Listener
}

func NewGrpcServerBuilder() InteractionServerBuilder {
	return &GrpcServerBuilder{}
}

func (b *GrpcServerBuilder) CreateServer(ctx context.Context) {
	b.server = grpc.NewServer()
}

func (b *GrpcServerBuilder) CreateListener(ctx context.Context, listenerConfig InteractionListenerConfig) error {
	var ntwType, ntwAddress string
	var ok bool

	if ntwType, ok = listenerConfig[NetworkType]; !ok {
		return errors.New("NETWORK TYPE CONFIG IS NOT PRESENTED")
	}

	if ntwAddress, ok = listenerConfig[ListenerAddress]; !ok {
		return errors.New("NETWORK ADDRESS CONFIG IS NOT PRESENTED")
	}

	var err error
	if b.listener, err = net.Listen(ntwType, ntwAddress); err != nil {
		return err
	}

	fmt.Printf("Listner will use [%s] network type on address [%s]", ntwType, ntwAddress)

	b.server = grpc.NewServer()
	return nil
}

func (b *GrpcServerBuilder) AddHooks(ctx context.Context, _map ServerHookMap) error {
	if b.server == nil {
		return errors.New("INTERACTION SERVER WAS NOT CONFIGURED")
	}

	var broken []string
	for k, v := range _map {
		if e := v(b.getServer); e != nil {
			broken = append(broken, k)
		}
	}

	if len(broken) > 0 {
		return errors.New("HOOK REGISTRATION PROCESS FAILED")
	}

	return nil

}

func (b *GrpcServerBuilder) GetResult() (serverBehavior InteractionServerBehavior, err error) {
	if b.server == nil {
		return nil, errors.New("INTERACTION SERVER WAS NOT CONFIGURED")
	} else if b.listener == nil {
		return nil, errors.New("INTERACTION SERVER LISTENER WAS NOT CONFIGURED")
	}

	return NewGrpcServer(b.server, b.listener), nil
}

func (b *GrpcServerBuilder) getServer() (server *grpc.Server) {
	return b.server
}
