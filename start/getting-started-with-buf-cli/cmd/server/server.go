package server

import (
	petv1 "Go_T_buffTutorial/gen/pet/v1"
	"Go_T_buffTutorial/internal/handler"
	"Go_T_buffTutorial/internal/pkg/gateway"
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
)

const (
	gRPCServerEndpoint = "localhost:30000"
	httpServerEndpoint = "localhost:50000"
)

type Server struct {
	gRPCGateway *gateway.GRPCGateway

	httpHandler http.Handler
	httpServer  *http.Server
	httpErrCh   chan error
	gRPCServer  *grpc.Server
	gRPCConn    *grpc.ClientConn
	gRPCErrCh   chan error

	PetStoreServiceHandler petv1.PetStoreServiceServer
}

type Option func(*Server) error

func WithGRPCGateway(g *gateway.GRPCGateway) Option {
	return func(s *Server) error {
		s.gRPCGateway = g
		return nil
	}
}

func NewServer(
	ctx context.Context,
	funcs []gateway.RegisterServiceHandlerFunc,
	opts ...Option,
) *Server {
	s := Server{}

	for _, v := range opts {
		err := v(&s)
		if err != nil {
			panic(err)
		}
	}

	s.gRPCServer = grpc.NewServer()
	s.PetStoreServiceHandler = handler.NewPetStoreServiceHandler()

	s.httpHandler = s.gRPCGateway.CreateHttpHandler(
		ctx,
		gRPCServerEndpoint,
		funcs...,
	)
	s.httpServer = &http.Server{Handler: s.httpHandler}
	return &s
}

func (s *Server) Start(ctx context.Context) error {
	petv1.RegisterPetStoreServiceServer(s.gRPCServer, s.PetStoreServiceHandler)

	ts, err := net.Listen("tcp", gRPCServerEndpoint)
	if err != nil {
		return err
	}
	hts, err := net.Listen("tcp", httpServerEndpoint)
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
