package main

import (
	petv1 "Go_T_buffTutorial/gen/pet/v1"
	"Go_T_buffTutorial/internal/handler"
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/handlers"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/protobuf/encoding/protojson"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const address = "localhost:8080"

type Server struct {
	httpHandler http.Handler
	httpServer  *http.Server
	httpErrCh   chan error
	gRPCServer  *grpc.Server
	gRPCConn    *grpc.ClientConn
	gRPCErrCh   chan error

	PetStoreServiceHandler petv1.PetStoreServiceServer
}

// func NewServer() *Server {
func NewServer(h http.Handler) *Server {
	s := Server{}
	s.gRPCServer = grpc.NewServer()
	s.PetStoreServiceHandler = handler.NewPetStoreServiceHandler()

	s.httpHandler = h
	s.httpServer = &http.Server{Handler: s.httpHandler}
	return &s
}

func (s *Server) Start(ctx context.Context) error {
	petv1.RegisterPetStoreServiceServer(s.gRPCServer, s.PetStoreServiceHandler)

	ts, err := net.Listen("tcp", fmt.Sprintf("localhost:30000"))
	if err != nil {
		return err
	}
	hts, err := net.Listen("tcp", fmt.Sprintf("localhost:50000"))
	if err != nil {
		// _ = ts.Close()
		return err
	}

	go func() {
		fmt.Println(">> httpServer.Serve")
		if err := s.httpServer.Serve(hts); !errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("HTTP Server error\n")
			//panic("HTTP Server error")
			s.httpErrCh <- err
		}
	}()
	go func() {
		fmt.Println(">> gRPCServer.Serve")
		if err := s.gRPCServer.Serve(ts); !errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("gRPC Server error\n")
			//panic("gRPC Server error")
			s.gRPCErrCh <- err
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP, syscall.SIGQUIT)
	select {
	case sig := <-ch:
		fmt.Printf("receive signal, sig: %s\n", sig.String())
	case err = <-s.httpErrCh:
		fmt.Printf("HTTP server stopped err %v\n", err)
	case err = <-s.gRPCErrCh:
		fmt.Printf("gRPC server stopped err %v\n", err)
	}
	return nil
}

// func (s *Server) GetHttpHandler(ctx context.Context, gRPCPort int) http.Handler {
func GetHttpHandler(ctx context.Context, gRPCPort int) http.Handler {
	endpoint := fmt.Sprintf("localhost:%d", gRPCPort)
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		panic(err)
	}
	//s.gRPCConn = conn

	mux := runtime.NewServeMux(
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

	_ = petv1.RegisterPetStoreServiceHandler(ctx, mux, conn)
	return handlers.CORS(
		handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "DELETE"}),
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedHeaders([]string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "Authorization", "ResponseType"}),
	)(mux)
}

func main() {
	ctx := context.Background()

	//s := NewServer()
	//s.httpHandler = GetHttpHandler(ctx, 30000)
	s := NewServer(GetHttpHandler(ctx, 30000))

	err := s.Start(ctx)
	if err != nil {
		fmt.Printf("main err: %#v\n", err)
	}
}
