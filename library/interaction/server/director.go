package server

import (
	"context"
	"errors"
)

type ServerBuildDirector struct {
	builder InteractionServerBuilder
}

func NewServerBuildDirector(builder InteractionServerBuilder) *ServerBuildDirector {
	if builder == nil {
		panic(errors.New("ILLEGAL ARGUMENT WAS PASSED: INTERACTION SERVER BUILDER"))
	}

	return &ServerBuildDirector{builder: builder}
}

func (b *ServerBuildDirector) CreatIneractionServer(ctx context.Context, config InteractionListenerConfig) (
	serverBehavior InteractionServerBehavior, err error) {

	b.builder.CreateServer(ctx)

	if myErr := b.builder.CreateListener(ctx, config); myErr == nil {
		return nil, myErr
	}

	var hookMap ServerHookMap
	if myErr := b.builder.AddHooks(ctx, hookMap); myErr == nil {
		return nil, myErr
	}

	return b.builder.GetResult()
}
