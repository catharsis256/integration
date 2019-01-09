package server

import (
	"context"
	"errors"
)

type ServerContextBehavior interface {
	ChangeState(newState ServerStateBehavior)
	GetServerBehavior() (serverBehavior InteractionServerBehavior, err error)
	SetServerBehavior(InteractionServerBehavior)
	GetListenerConfig() (InteractionListenerConfig, error)
}

type ServerContext struct {
	serverBehavior InteractionServerBehavior
	stateBehavior  ServerStateBehavior
	listenerConfig InteractionListenerConfig
}

func NewServerContext(interactionServBehavior InteractionServerBehavior) *ServerContext {
	if interactionServBehavior == nil {
		panic(errors.New("ILLEGAL ARGUMENT WAS PASSED: INTERACTION SERVER BEHAVIOR"))
	}

	return &ServerContext{stateBehavior: NewUninitializedState()}
}

func (c *ServerContext) ChangeState(newState ServerStateBehavior) {
	c.stateBehavior = newState
}

func (c *ServerContext) GetServerBehavior() (serverBehavior InteractionServerBehavior, err error) {
	if c.serverBehavior != nil {
		return c.serverBehavior, nil
	}

	return nil, errors.New("INTERACTION SERVER BEHAVIOR WAS NOT CONFIGURED")
}

func (c *ServerContext) SetServerBehavior(serverBehavior InteractionServerBehavior) {
	if serverBehavior == nil {
		panic(errors.New("INTERACTION SERVER BEHAVIOR WAS NOT CONFIGURED"))
	}

	c.serverBehavior = serverBehavior
}

func (c *ServerContext) GetListenerConfig() (listenerConfig InteractionListenerConfig, err error) {
	if c.listenerConfig == nil {
		panic(errors.New("INTERACTION SERVER CONFIGURATION DID NOT SET"))
	}

	return c.listenerConfig, nil
}

//-----------------------------------
type ServerStateBehavior interface {
	Inquire(context.Context, ServerContextBehavior) error
}

//-----------------------------------
type UninitializedState struct{}

func NewUninitializedState() ServerStateBehavior {
	return &UninitializedState{}
}

func (s *UninitializedState) Inquire(ctx context.Context, context ServerContextBehavior) error {
	buildDirector := NewServerBuildDirector(NewGrpcServerBuilder())

	var err error
	var config InteractionListenerConfig
	if config, err = context.GetListenerConfig(); err != nil {
		return err
	}

	var serverBehavior InteractionServerBehavior
	if serverBehavior, err = buildDirector.CreatIneractionServer(ctx, config); err != nil {
		return err
	}

	context.SetServerBehavior(serverBehavior)

	return nil
}

//-----------------------------------
type InitiateServingState struct{}

func NewInitiateServingState() ServerStateBehavior {
	return &InitiateServingState{}
}

func (s *InitiateServingState) Inquire(ctx context.Context, context ServerContextBehavior) error {
	var servBehavior InteractionServerBehavior
	var err error

	if servBehavior, err = context.GetServerBehavior(); err != nil {
		return err
	}

	if err = servBehavior.Serve(); err != nil {
		return err
	}

	context.ChangeState(NewHandlingRequestState())
	return nil
}

//-----------------------------------
type RequestHandlingState struct {
}

func NewHandlingRequestState() ServerStateBehavior {
	return &RequestHandlingState{}
}

func (s *RequestHandlingState) Inquire(ctx context.Context, context ServerContextBehavior) error {
	// stub

	return nil
}
