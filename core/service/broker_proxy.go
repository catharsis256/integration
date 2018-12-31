package service

import (
	"com.smartblocklab/integration/core/model"
	"context"
	"errors"
)

type SupervisorProxy struct {
	Service SupervisorInterface
}

func NewSupervisorProxy(supervisor SupervisorInterface) ProxyInterface {
	return &SupervisorProxy{Service: supervisor}
}

func (h *SupervisorProxy) GetService() interface{} {
	return h.Service
}

func (h *SupervisorProxy) Initiate(ctx context.Context) {
	serv := h.GetService()

	if sInterface, ok := serv.(SupervisorInterface); ok {
		sInterface.Initiate(ctx)
		println("SupervisorInterface initiated")
		return
	}

	panic(errors.New(model.ErrorNotSuitableProxy))
}