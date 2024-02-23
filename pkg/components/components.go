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

	SubConfigurator[SubConfig any] interface {
		Start(ctx context.Context) error
		Stop(ctx context.Context) error
		Configure(ctx context.Context, config SubConfig) error
	}

	ComponentFunc[Config any] func(ctx context.Context, config Config) error

	Component[Config any] struct {
		configurator SubConfigurator[Config]
		status       Status
	}

	Components[Config any] struct {
		Components []*Component[Config]
	}
)

type ConfigureFunc[Config any] func(ctx context.Context, config Config, init bool) error

type ConfigAdapter[Config any, SubConfig any] func(self Config) SubConfig

func (c *Components[Config]) Add(configurator Configurator[Config]) {
	component := &Component[Config]{
		configurator: configurator,
		status:       Stopped,
	}

	c.Components = append(c.Components, component)

}

func AddComponent[
	Config any,
	SubConfig any,
	Conf Configurator[SubConfig],
](
	components Components[Config],
	configurator Conf,
	adapter ConfigAdapter[Config, SubConfig],
) {
	components.Add(configurator)
	//components.Add(func(ctx context.Context, config Config, init bool) error {
	//	return module.Configure(ctx, adapter(config), init)
	//})
}

func Add[Config any, SubConfig any, Group intoGroup[Config], Module Configurator[SubConfig]](
	group Group,
	module Module,
	adapter ConfigAdapter[Config, SubConfig],
) {
	group.intoGroup().AddModule(module, func(ctx context.Context, config Config, init bool) error {
		return module.Configure(ctx, adapter(config), init)
	})
}
