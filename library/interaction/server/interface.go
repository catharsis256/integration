package server

import (
	"context"
	"google.golang.org/grpc"
)

const (
	NetworkType     = "NetworkType"
	ListenerAddress = "ListenerAddress"
)

type ServerProducer func() *grpc.Server
type ServerHookFunction func(serverProducer ServerProducer) error
type ServerHookMap map[string]ServerHookFunction
type InteractionListenerConfig map[string]string

type InteractionServerBuilder interface {
	CreateServer(context.Context)
	CreateListener(context.Context, InteractionListenerConfig) error
	AddHooks(context.Context, ServerHookMap) error
	GetResult() (InteractionServerBehavior, error)
}

type InteractionServerBehavior interface {
	Serve() error
	Stop()
}
