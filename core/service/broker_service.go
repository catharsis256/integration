package service

type BrokerServiceInterface interface {
	InstantiateService
}

type AbstractBrokerService struct {

}

type SimpleBrokerService struct {
	AbstractBrokerService
}

func (s *SimpleBrokerService) instance(InputParameterMap) error {
	return nil
}
