package main

import (
	"context"
	"fmt"
	"gitlab.smartblocklab.com/integration/library/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":4444", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := conn.Close(); err != nil {
			panic(err)
		}
	}()

	client := proto.NewGreaterServiceClient(conn)
	ctx := context.Background()
	request := proto.GreeterRequest{Greeting: "Hello", Name: "Gopher"}

	response, err := client.Greet(ctx, &request)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", response)

}
