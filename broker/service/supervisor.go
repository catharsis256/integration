package service

import (
	"context"
	"gitlab.smartblocklab.com/integration/library/interfaces"
)

type SupervisorInterface interface {
	interfaces.LaunchInterface
}

func NewSupervisorService() SupervisorInterface {
	return &SupervisorService{}
}

type SupervisorService struct {
}

func (b *SupervisorService) Launch(ctx context.Context) error {
	println("SupervisorService launched")
	return nil
}
