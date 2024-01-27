package grpc

import (
	"context"
	pb "mir-service/internal/server/grpc-gateway/pb"
)

//type Client[T, R any] interface {
//	Do(ctx context.Context, value T) (R, error)
//}
//
//type fromCore[T, R any] interface {
//	FromCore(value T) R
//}

//type Client2 interface {
//	Do(ctx context.Context, loginUserRequest *pb.LoginUserRequest) (*LoginUserRequest, error)
//}

//func FromCore[T any](ctx context.Context, value T) error {
//	value.Do()
//	return nil
//}

//func FromCore[T, R Client[T, R]](value T) {
//	value.Do()
//}

func (s *Server) Login(ctx context.Context, in *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	return in.Do(ctx, s.core)
}

//func LoginUserRequest(ctx context.Context,
//	in *pb.LoginUserRequest) (*pb.LoginUserRequest, error) {
// Сюда добавить вызов функции Do
//loginUserRequest := LoginUserRequest{Username: "te", Password: "1"}
//loginUserRequest := LoginUserRequest{Username: "te", Password: "1"}
//_, err := loginUserRequest.Do(ctx, in)
//_, err := Client.Do(ctx, in)
////FromCore(ctx, in)
//if err != nil {
//	return nil, nil
//}
//	fromCore.FromCore(in)
//	return nil, nil
//}

func (r *pb.LoginUserRequest) Do(ctx context.Context, c *Core) (*pb.LoginUserResponse, error) {
	//return callCore(ctx, c, r.ToCore().Do, new(GetGroupsResponse))
	return nil, nil
}

//func (l *pb.LoginUserRequest) ToCore(ctx context.Context) (*LoginUserRequest, error) {
//
//	return FromCore(ctx, l, new(LoginUserRequest))
//}

//	type fromCore[T, R any] interface {
//		FromCore(value T) R
//	}
//type toCore[T any, R any] interface {
//	ToCore(ctx context.Context, value T) (R, error)
//}

//func ToCore(ctx context.Context, loginUserRequest *pb.LoginUserRequest) (*LoginUserRequest, error) {
//	return FromCore(ctx, loginUserRequest, new(LoginUserRequest))
//}
//
//func FromCore(ctx context.Context, loginUserRequest *pb.LoginUserRequest, loginUserLocalRequest *LoginUserRequest) (*LoginUserRequest, error) {
//	return loginUserLocalRequest.Do(ctx, loginUserRequest)
//}

//type LoginUserRequest1 struct {
//	Username string
//	Password string
//}
//
//func (l *LoginUserRequest1) FromCore(loginUserRequest *pb.LoginUserRequest) *LoginUserRequest1 {
//	l.Username = loginUserRequest.Username
//	l.Password = loginUserRequest.Password
//	return nil
//}
//
//func (l *LoginUserRequest1) Do(ctx context.Context, loginUserRequest *pb.LoginUserRequest) (*LoginUserRequest1, error) {
//	l.Username = loginUserRequest.Username
//	l.Password = loginUserRequest.Password
//	return nil, nil
//}
