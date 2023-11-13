package server

import (
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

type Manager struct {
	gRPCGateway *gateway.GRPCGateway

	httpHandler http.Handler
	httpServer  *http.Server
	httpErrCh   chan error
	gRPCServer  *grpc.Server
	gRPCConn    *grpc.ClientConn
	gRPCErrCh   chan error
}

type Option func(*Manager) error

func WithGRPCGateway(g *gateway.GRPCGateway) Option {
	return func(s *Manager) error {
		s.gRPCGateway = g
		return nil
	}
}

func NewManager(
	ctx context.Context,
	registerServiceHandlerFuncs []gateway.RegisterServiceHandlerFunc,
	registerServiceServerFuncs []gateway.RegisterServiceServerFunc,
	opts ...Option,
) *Manager {
	m := Manager{}

	for _, v := range opts {
		err := v(&m)
		if err != nil {
			panic(err)
		}
	}

	m.gRPCServer = grpc.NewServer()
	for _, v := range registerServiceServerFuncs {
		v(ctx, m.gRPCServer)
	}

	m.httpHandler = m.gRPCGateway.CreateHttpHandler(
		ctx,
		gRPCServerEndpoint,
		registerServiceHandlerFuncs...,
	)
	m.httpServer = &http.Server{Handler: m.httpHandler}
	return &m
}

func (m *Manager) Start(_ context.Context) error {
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
		fmt.Printf(">> httpServer.Serve: %s\n", gRPCServerEndpoint)
		if err := m.httpServer.Serve(hts); !errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("HTTP Server error\n")
			m.httpErrCh <- err
		}
	}()
	go func() {
		fmt.Printf(">> gRPCServer.Serve: %s\n", httpServerEndpoint)
		if err := m.gRPCServer.Serve(ts); !errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("gRPC Server error\n")
			m.gRPCErrCh <- err
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP, syscall.SIGQUIT)
	select {
	case sig := <-ch:
		fmt.Printf("receive signal, sig: %s\n", sig.String())
	case err = <-m.httpErrCh:
		fmt.Printf("HTTP server stopped err %v\n", err)
	case err = <-m.gRPCErrCh:
		fmt.Printf("gRPC server stopped err %v\n", err)
	}
	return nil
}
