package main

import (
	"Go_T_buffTutorial/internal/pkg/gateway"
	"Go_T_buffTutorial/internal/pkg/server"
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()

	s := server.NewManager(
		ctx,
		RegisterServiceHandlerFuncs(),
		RegisterServiceServerFuncs(),
		server.WithGRPCGateway(gateway.NewGRPCGateway()),
	)

	err := s.Start(ctx)
	if err != nil {
		fmt.Printf("main err: %#v\n", err)
	}
}
