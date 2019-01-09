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

func (b *ServerBuildDirector) CreatIneractionServer(
					ctx context.Context,
					config InteractionListenerConfig) (serverBehavior InteractionServerBehavior, _err error) {
	b.builder.CreateServer(ctx)

	if err := b.builder.CreateListener(ctx, config); err == nil {
		return nil, err
	}

	var hookMap ServerHookMap
	if err := b.builder.AddHooks(ctx, hookMap); err == nil {
		return nil, nil
	}

	return b.builder.GetResult()
}