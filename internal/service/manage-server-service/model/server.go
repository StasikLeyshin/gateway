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

	GetServerResponse struct {
		Servers []Server
	}
)
