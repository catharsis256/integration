package service

type SupervisorHolder struct {
	Service SupervisorInterface
}

func NewSupervisorHolder(supervisor SupervisorInterface) SupervisorHolderInterface {
	return &SupervisorHolder{Service: supervisor}
}

func (h *SupervisorHolder) GetService() SupervisorInterface {
	return h.Service
}

type SupervisorHolderInterface interface {
	GetService() SupervisorInterface
}
