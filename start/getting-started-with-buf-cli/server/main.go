package main

import (
	petv1 "Go_T_buffTutorial/gen/pet/v1"
	"Go_T_buffTutorial/internal/handler"
	"context"
	"fmt"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/protobuf/encoding/protojson"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const address = "localhost:8080"

type Server struct {
	gRPCServer             *grpc.Server
	gRPCConn               *grpc.ClientConn
	PetStoreServiceHandler petv1.PetStoreServiceServer
}

func NewServer() *Server {
	s := Server{}
	s.gRPCServer = grpc.NewServer()
	s.PetStoreServiceHandler = handler.NewPetStoreServiceHandler()
	return &s
}

func (s *Server) Start(ctx context.Context) error {
	petv1.RegisterPetStoreServiceServer(s.gRPCServer, s.PetStoreServiceHandler)
	return nil
}

func (s *Server) GetHttpHandler(ctx context.Context, gRPCPort int) http.Handler {
	endpoint := fmt.Sprintf("127.0.0.1:%d", gRPCPort)
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		panic(err)
	}
	s.gRPCConn = conn

	runtime.NewServeMux(
		runtime.WithMarshalerOption(
			runtime.MIMEWildcard,
			&runtime.HTTPBodyMarshaler{
				Marshaler: &runtime.JSONPb{
					MarshalOptions: protojson.MarshalOptions{
						EmitUnpopulated: true,
						UseEnumNumbers:  true,
					},
					UnmarshalOptions: protojson.UnmarshalOptions{
						DiscardUnknown: true,
					},
				},
			},
		),
	)

	_ = petv1.RegisterPetStoreServiceServer()
}

func main() {
	ctx := context.Background()
	s := NewServer()

	err := s.Start(ctx)
	if err != nil {
		fmt.Errorf("main err: %#v\n", err)
	}
}
