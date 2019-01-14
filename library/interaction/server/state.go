package server

import "golang.org/x/net/context"

type ServerStateBehavior interface {
	Inquire(context.Context, ServerContextBehavior) error
}

type BaseState struct {}

type UninitializedState struct{
	BaseState
}
type InitiateServingState struct{
	BaseState
}
type RequestHandlingState struct {
	BaseState
}
type ShutdownState struct {
	BaseState
}
type FinalizedState struct {
	BaseState
}


func NewUninitializedState() ServerStateBehavior {
	return &UninitializedState{}
}

func NewInitiateServingState() ServerStateBehavior {
	return &InitiateServingState{}
}

func NewHandlingRequestState() ServerStateBehavior {
	return &RequestHandlingState{}
}

func NewShutdownState() ServerStateBehavior {
	return &ShutdownState{}
}

func NewFinalizedState() ServerStateBehavior {
	return &FinalizedState{}
}



func (s *BaseState) Inquire(ctx context.Context, context ServerContextBehavior) error {
	return nil
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

	context.ChangeState(NewInitiateServingState())

	return nil
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


func (s *ShutdownState) Inquire(ctx context.Context, context ServerContextBehavior) error {
	var servBehavior InteractionServerBehavior
	var err error


	if servBehavior, err = context.GetServerBehavior(); err != nil {
		return err
	}

	servBehavior.Stop()

	context.ChangeState(NewFinalizedState())
	return nil
}

