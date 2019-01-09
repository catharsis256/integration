package server

import "errors"

type ServerContextBehavior interface {
	ChangeState(newState ServerStateBehavior)
	GetServerBehavior() (serverBehavior InteractionServerBehavior, err error)
}

type ServerContext struct {
	serverBehavior InteractionServerBehavior
	stateBehavior ServerStateBehavior
}

func NewServerContext(interactionServBehavior InteractionServerBehavior) *ServerContext {
	if interactionServBehavior == nil {
		panic(errors.New("ILLEGAL ARGUMENT WAS PASSED: INTERACTION SERVER BEHAVIOR"))
	}

	return &ServerContext{stateBehavior: NewInitiateServingState(), serverBehavior: interactionServBehavior}
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


//-----------------------------------
type ServerStateBehavior interface {
	Inquire(ServerContextBehavior) error
}

//-----------------------------------
type UninitializedState struct { }

func NewUninitializedState() ServerStateBehavior {
	return &UninitializedState{}
}

func (s *UninitializedState) Inquire(context ServerContextBehavior) error {
	// stub
	buildDirector := NewServerBuildDirector(NewGrpcServerBuilder())
	return nil
}

//-----------------------------------
type InitiateServingState struct { }

func NewInitiateServingState() ServerStateBehavior {
	return &InitiateServingState{}
}

func (s *InitiateServingState) Inquire(context ServerContextBehavior) error {
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

func (s *RequestHandlingState) Inquire(context ServerContextBehavior) error {
	// stub

	return nil
}
