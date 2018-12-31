package service

import "context"

type InitiateInterface interface {
	Initiate(context.Context)
}

type ProxyInterface interface {
	GetService() interface{}
}
