package interfaces

import "context"

type InitiateDefaultInterface interface {
	Initiate(context.Context) error
}

type LaunchInterface interface {
	Launch(ctx context.Context) error
}
