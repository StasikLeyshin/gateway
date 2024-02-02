package grpc

//import (
//	"fmt"
//	pb "gateway/internal/server/grpc-gateway/pb"
//	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
//	"golang.org/x/net/context"
//	"google.golang.org/grpc"
//	"google.golang.org/grpc/credentials/insecure"
//	"google.golang.org/protobuf/encoding/protojson"
//	"net"
//	"net/http"
//)
//
//type Config struct {
//	Port  int  `yaml:"port"`
//	Debug bool `yaml:"debug"`
//}
//
//type Core struct {
//}
//
//type Server struct {
//	pb.UnsafeGatewayServer
//	//client Client
//	server *http.Server
//	core   Core
//	debug  bool
//	port   int
//}
//
//func NewServer(config Config) *Server {
//	//server, err := gapi.NewServer(config, store, taskDistributor)
//	//if err != nil {
//	//	log.Fatal().Err(err).Msg("cannot create server")
//	//}
//
//	server := Server{debug: true, port: config.Port}
//
//	jsonOption := runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
//		MarshalOptions: protojson.MarshalOptions{
//			UseProtoNames: true,
//		},
//		UnmarshalOptions: protojson.UnmarshalOptions{
//			DiscardUnknown: true,
//		},
//	})
//
//	grpcMux := runtime.NewServeMux(jsonOption)
//
//	ctx, cancel := context.WithCancel(context.Background())
//	defer cancel()
//
//	//err := pb.RegisterGatewayHandlerServer(ctx, grpcMux, server)
//	//err := pb.RegisterGatewayServer(grpcMux, server)
//	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
//	err := pb.RegisterGatewayHandlerFromEndpoint(ctx, grpcMux, "localhost:12201", opts)
//	if err != nil {
//		//log.Fatal().Err(err).Msg("cannot register handler server")
//	}
//
//	mux := http.NewServeMux()
//	mux.Handle("/", grpcMux)
//
//	//statikFS, err := fs.New()
//	//if err != nil {
//	//	log.Fatal().Err(err).Msg("cannot create statik fs")
//	//}
//
//	//swaggerHandler := http.StripPrefix("/swagger/", http.FileServer(statikFS))
//	//mux.Handle("/swagger/", swaggerHandler)
//
//	listener, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:", config.Port))
//	if err != nil {
//		//log.Fatal().Err(err).Msg("cannot create listener")
//	}
//
//	//log.Info().Msgf("start HTTP gateway server at %s", listener.Addr().String())
//	handler := HttpLogger(mux)
//	err = http.Serve(listener, handler)
//	if err != nil {
//		//log.Fatal().Err(err).Msg("cannot start HTTP gateway server")
//	}
//	return &server
//}
//
//func (s *Server) Start() {
//
//}
//
//func (s *Server) Stop() {
//
//}
