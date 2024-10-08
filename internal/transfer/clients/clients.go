package clients

import (
	"context"
	descInternal "github.com/StasikLeyshin/libs-proto/grpc/role-service/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
	"net/url"
)

type GRPCClient struct {
	descInternal.RoleServiceClient
	*HTTPClient
}

func NewGRPCClient(ctx context.Context, address string) *GRPCClient {
	options := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	connection, _ := grpc.Dial(address, options...)

	return &GRPCClient{
		RoleServiceClient: descInternal.NewRoleServiceClient(connection),
		HTTPClient:        newHTTPClient(address),
	}
}

type HTTPClient struct {
	route  string
	client *http.Client
}

func newHTTPClient(route string) *HTTPClient {
	return &HTTPClient{
		route:  route,
		client: http.DefaultClient,
	}
}

func (c *HTTPClient) AddHostAddress(address url.URL) (url.URL, error) {
	address.Host = c.route
	address.Scheme = "http"

	return address, nil
}
