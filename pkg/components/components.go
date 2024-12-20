package components

import (
	"context"
	"fmt"
)

const (
	Stopped Status = 0
	Started Status = 1
)

var ComponentError = map[Status]string{
	Started: "This component has already started",
	Stopped: "This component has already been stopped",
}

type (
	Status int

	Configurator[Config any] interface {
		Start(ctx context.Context) error
		Stop(ctx context.Context) error
		Configure(ctx context.Context, config Config) error
	}

	Control[SubConfig any] interface {
		Start(ctx context.Context) error
		Stop(ctx context.Context) error
	}

	ComponentFunc[Config any]                func(ctx context.Context, config Config) error
	ConfigAdapter[Config any, SubConfig any] func(self Config) SubConfig

	Component[Config any] struct {
		control       Control[Config]
		configuration ComponentFunc[Config]
		name          string
		status        Status
	}

	Components[Config any] struct {
		components []*Component[Config]
	}
)

func NewComponents[Config any]() *Components[Config] {
	return &Components[Config]{}
}

func (c *Components[Config]) Add(configurator Control[Config], configuration ComponentFunc[Config], name string) {
	component := &Component[Config]{
		control:       configurator,
		configuration: configuration,
		name:          name,
		status:        Stopped,
	}

	c.components = append(c.components, component)
}

func (c *Components[Config]) actionError(component *Component[Config], text string, err error) error {
	return fmt.Errorf("component name: %s | description: %s | error: %v", component.name, text, err)
}

func (c *Components[Config]) checkStatus(component *Component[Config], status Status) error {
	if component.status != status {
		component.status = status
	} else {
		return c.actionError(component, ComponentError[status], nil)
	}
	return nil
}

func (c *Components[Config]) Start(ctx context.Context) []error {
	errors := make([]error, 0, len(c.components))
	for _, component := range c.components {
		err := component.control.Start(ctx)
		if err != nil {
			errors = append(errors, c.actionError(component, "error when starting the component", err))
		} else {
			if err = c.checkStatus(component, Started); err != nil {
				errors = append(errors, err)
			}
		}
	}
	return errors
}

func (c *Components[Config]) Stop(ctx context.Context) []error {
	errors := make([]error, 0, len(c.components))
	for _, component := range c.components {
		err := component.control.Stop(ctx)
		if err != nil {
			errors = append(errors, c.actionError(component, "error when stopping the component", err))
		} else {
			if err = c.checkStatus(component, Stopped); err != nil {
				errors = append(errors, err)
			}
		}
	}
	return errors
}

func (c *Components[Config]) Configure(ctx context.Context, config Config) []error {
	errors := make([]error, 0, len(c.components))
	for _, component := range c.components {
		err := component.configuration(ctx, config)
		if err != nil {
			errors = append(errors, c.actionError(component, "component configuration error", err))
		}
	}
	return errors
}

func AddComponent[
	Config any,
	SubConfig any,
	Conf Configurator[SubConfig],
](
	components *Components[Config],
	configurator Conf,
	adapter ConfigAdapter[Config, SubConfig],
	name string,
) {
	components.Add(configurator, func(ctx context.Context, config Config) error {
		return configurator.Configure(ctx, adapter(config))
	}, name)
}
