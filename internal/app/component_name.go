package app

type ComponentName string

const (
	ComponentNameGRPCServer      ComponentName = "GRPC_Server"
	ComponentNameConnector       ComponentName = "Connector"
	ComponentNameService         ComponentName = "Service"
	ComponentNameServiceProvider ComponentName = "Provider"

	ComponentNameRepositoryMongo ComponentName = "MongoDB"
)

func (name ComponentName) String() string {
	return string(name)
}
