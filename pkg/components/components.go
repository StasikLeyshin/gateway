package components

import "context"

const (
	Stopped Status = 0
	Started Status = 1
)

type (
	//Starter interface {
	//	Start() error
	//}
	//
	//Stopper interface {
	//	Stop(ctx context.Context) error
	//}
	//
	//Configurator[Config any] interface {
	//	Configure(ctx context.Context, config Config) error
	//}

	Status int

	Configurator[Config any] interface {
		Start(ctx context.Context) error
		Stop(ctx context.Context) error
		Configure(ctx context.Context, config Config) error
	}

	ComponentFunc[Config any] func(ctx context.Context, config Config) error

	Component[Config any] struct {
		configurator Configurator[Config]
		status       Status
	}

	Components[Config any] struct {
		Components []Component[Config]
	}
)

func (c *Components[Config]) Add() {

}

func AddComponent[Config any](configurator Configurator[Config]) {

}
