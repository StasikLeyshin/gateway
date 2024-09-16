package role

import (
	"context"
	"gateway/internal/transfer/grpc/role/model"
	descInternal "github.com/StasikLeyshin/libs-proto/grpc/role-service/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
	"net/url"
)

type Client struct {
	descInternal.RoleServiceClient
	*HTTPClient
}

func NewClient(ctx context.Context, address string) *Client {
	options := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	connection, _ := grpc.Dial(address, options...)

	return &Client{
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

func (i *transferRoleService) Login(ctx context.Context, request *model.LoginRequest) (*model.LoginResponse, error) {
	NewClient(ctx, "").RoleServiceClient.Login(ctx, request.FromTransfer())
	return nil, nil
}
