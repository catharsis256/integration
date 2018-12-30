package service

type ProgramServices uint8
type InputServiceParameters uint16
type InputParameterMap map[InputServiceParameters]interface{}

const (
	SimpleBrokerServiceName ProgramServices = iota
)

const (
	_ InputServiceParameters = iota
)

type InstantiateService interface {
	instance(InputParameterMap) error
}

func Instantiate(ps ProgramServices, params InputParameterMap) (instantiatedService chan<- InstantiateService, error error) {

	switch ps {

	case SimpleBrokerServiceName:
		go setupInstance(instantiatedService, &SimpleBrokerService{}, params)

	}

	return
}

func setupInstance(serviceChannel chan<- InstantiateService, service InstantiateService, params InputParameterMap) {
	if err := service.instance(params); err != nil {
		panic(err)
	}
	serviceChannel <- service
}
