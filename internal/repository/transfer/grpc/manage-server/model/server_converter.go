package model

import desc "github.com/StasikLeyshin/libs-proto/grpc/manage-server-service/pb"

func (convert *Server) ToTransfer(value *desc.Server) *Server {
	request := &Server{
		ServerID:   value.ServerId,
		NodeID:     value.NodeId,
		ServerType: new(ServerType).ToTransfer(value.ServerType),
		Name:       value.Name,
		Host:       value.Host,
		Port:       value.Port,
	}

	return request
}

var (
	serverTypeToTransfer = map[desc.ServerType]ServerType{
		desc.ServerType(0): GatewayServerType,
		desc.ServerType(1): RoleServerType,
		desc.ServerType(2): ManageServerServerType,
	}

	serverTypeFromTransfer = map[ServerType]desc.ServerType{
		GatewayServerType:      desc.ServerType(0),
		RoleServerType:         desc.ServerType(1),
		ManageServerServerType: desc.ServerType(2),
	}
)

func (convert ServerType) ToTransfer(value desc.ServerType) ServerType {
	result, _ := serverTypeToTransfer[value]
	return result
}

func (convert ServerType) FromTransfer() desc.ServerType {
	result, _ := serverTypeFromTransfer[convert]
	return result
}

// GetServers

func (convert *GetServersRequest) FromTransfer() *desc.GetServersRequest {
	if convert == nil {
		return nil
	}

	request := &desc.GetServersRequest{}

	request.NodeId = convert.NodeID
	request.Name = convert.Name
	request.Host = convert.Host
	request.Port = convert.Port

	if convert.ServerType != nil {
		request.ServerType = make([]desc.ServerType, 0, len(convert.ServerType))
	}

	for i := range convert.ServerType {
		request.ServerType = append(request.ServerType, convert.ServerType[i].FromTransfer())
	}

	return request
}

func (convert *GetServersResponse) ToTransfer(value *desc.GetServersResponse) *GetServersResponse {
	if value.Servers != nil {
		convert.Servers = make([]*Server, 0, len(value.Servers))
	}

	for i := range value.Servers {
		convert.Servers = append(convert.Servers, new(Server).ToTransfer(value.Servers[i]))
	}

	return convert
}

// GetServersAddresses

func (convert *GetServersAddressesRequest) FromTransfer() *desc.GetServersAddressesRequest {
	if convert == nil {
		return nil
	}

	request := &desc.GetServersAddressesRequest{}

	request.ServerType = convert.ServerType.FromTransfer()

	return request
}

func (convert *GetServersAddressesResponse) ToTransfer(value *desc.GetServersAddressesResponse) *GetServersAddressesResponse {
	if value.Servers != nil {
		convert.Servers = make([]*Server, 0, len(value.Servers))
	}

	for i := range value.Servers {
		convert.Servers = append(convert.Servers, new(Server).ToTransfer(value.Servers[i]))
	}

	return convert
}
