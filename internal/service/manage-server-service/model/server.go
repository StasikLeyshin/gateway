package model

type (
	Server struct {
		ID         string
		Name       string
		Host       string
		Port       string
		ServerType string
	}
)

type (
	GetServersRequest struct {
		ServerType string
		Name       *string
		Host       *string
		Port       *string
	}

	GetServersResponse struct {
		Servers []Server
	}
)
