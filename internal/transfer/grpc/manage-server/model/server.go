package model

type (
	Server struct {
		ServerID   string
		NodeID     string
		Name       string
		Host       string
		Port       string
		ServerType ServerType
	}

	ServerType string
)

const (
	GatewayServerType      = "server_type_gateway"
	RoleServerType         = "server_type_role"
	ManageServerServerType = "server_type_manage_server"
)

type (
	GetServersRequest struct {
		ServerType []ServerType
		NodeID     *string
		Name       *string
		Host       *string
		Port       *string
	}

	GetServersResponse struct {
		Servers []*Server
	}
)

type (
	GetServersAddressesRequest struct {
		ServerType ServerType
	}

	GetServersAddressesResponse struct {
		Servers []*Server
	}
)
