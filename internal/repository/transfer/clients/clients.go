package clients

import (
	"context"
	server "github.com/StasikLeyshin/libs-proto/grpc/manage-server-service/pb"
	role "github.com/StasikLeyshin/libs-proto/grpc/role-service/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	role.RoleServiceClient
	server.ManageServiceClient
	*HTTPClient
}

func NewClient(ctx context.Context, address string) *Client {
	options := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	connection, _ := grpc.Dial(address, options...)

	return &Client{
		RoleServiceClient:   role.NewRoleServiceClient(connection),
		ManageServiceClient: server.NewManageServiceClient(connection),
		HTTPClient:          newHTTPClient(address),
	}
}

type HTTPClient struct {
	route  string
	client *http.Client
}

func newHTTPClient(route string) *HTTPClient {
	return &HTTPClient{
		route: route,
		client: &http.Client{
			Timeout: time.Second * 10, // TODO: возможно вынести в конфиг
		},
	}
}

func (c *HTTPClient) AddHostAddress(address url.URL) (url.URL, error) {
	address.Host = c.route
	address.Scheme = "http"

	return address, nil
}
