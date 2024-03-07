module gateway

go 1.21.4

require (
	github.com/StasikLeyshin/libs-proto/grpc-gateway/role-service v0.0.7
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.18.1
	github.com/sirupsen/logrus v1.9.3
	google.golang.org/genproto/googleapis/api v0.0.0-20231106174013-bbf56f31fb17
	google.golang.org/grpc v1.59.0
	google.golang.org/protobuf v1.31.0
	gopkg.in/yaml.v3 v3.0.1
)

require github.com/StasikLeyshin/libs-proto/grpc/manage-server-service v0.0.10

require github.com/matoous/go-nanoid/v2 v2.0.0 // indirect

require (
	github.com/golang/protobuf v1.5.3 // indirect; indirec
	github.com/kr/text v0.2.0 // indirect
	github.com/rogpeppe/go-internal v1.11.0 // indirect
	golang.org/x/net v0.17.0
	golang.org/x/sys v0.13.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto v0.0.0-20231030173426-d783a09b4405 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20231030173426-d783a09b4405 // indirect
)
