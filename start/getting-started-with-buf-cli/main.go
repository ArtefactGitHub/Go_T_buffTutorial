package main

import (
	"Go_T_buffTutorial/server"
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()

	s := server.NewServer(server.GetHttpHandler(ctx, 30000))

	err := s.Start(ctx)
	if err != nil {
		fmt.Printf("main err: %#v\n", err)
	}
}
