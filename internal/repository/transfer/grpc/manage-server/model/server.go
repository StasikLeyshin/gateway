package model

import (
	desc "github.com/StasikLeyshin/libs-proto/grpc/manage-server-service/pb"
	"github.com/golang/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoimpl"
	"google.golang.org/protobuf/types/descriptorpb"
)

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
	//ServerName string
)

const (
	GatewayServerType      ServerType = "server_type_gateway"
	RoleServerType         ServerType = "server_type_role"
	ManageServerServerType ServerType = "server_type_manage_server"

	//ServerName = desc.E_ServiceName. proto. //proto.GetExtension(md, desc.E_ServiceName)
)

var ServerName *string

func init() {
	ServerName = GetServerName(
		desc.File_grpc_manage_server_service_proto_manage_server_service_proto,
		desc.E_ServiceName,
	)
}

func GetServerName(fileDesc protoreflect.FileDescriptor, serviceName *protoimpl.ExtensionInfo) *string {
	opts, ok := fileDesc.Options().(*descriptorpb.FileOptions)
	if !ok {
		return nil
	}

	serverName, _ := proto.GetExtension(opts, serviceName)

	return serverName.(*string)

}

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