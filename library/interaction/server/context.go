package server

import (
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