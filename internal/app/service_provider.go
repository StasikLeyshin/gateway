package app

import "gateway/internal/app/startup"

type serviceProvider struct {
	config *startup.Config
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}
