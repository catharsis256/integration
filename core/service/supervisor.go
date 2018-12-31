package service

import "context"


type SupervisorInterface interface {
	InitiateInterface
}

func NewSupervisorService() SupervisorInterface {
	return &SupervisorService{}
}

type SupervisorService struct {

}

func (b *SupervisorService) Initiate(ctx context.Context) {
	println("SupervisorService initiated")
}

