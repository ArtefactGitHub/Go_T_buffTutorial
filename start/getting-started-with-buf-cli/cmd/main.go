package main

import (
	"Go_T_buffTutorial/cmd/server"
	"Go_T_buffTutorial/internal/pkg/gateway"
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()

	s := server.NewServer(
		ctx,
		server.RegisterServiceHandlerFuncs(),
		server.WithGRPCGateway(gateway.NewGRPCGateway()),
	)

	err := s.Start(ctx)
	if err != nil {
		fmt.Printf("main err: %#v\n", err)
	}
}
