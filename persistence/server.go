package persistence

import (
	"fmt"
	"gitlab.smartblocklab.com/integration/library/proto"
	"gitlab.smartblocklab.com/integration/persistence/server"
	"google.golang.org/grpc"
	"net"
)

func main() {
	grpcServer := grpc.NewServer()

	grpcServer.GetServiceInfo()
	proto.RegisterGreaterServiceServer(grpcServer, &server.Greeter{Exclaim: true})

	listen, err := net.Listen("tcp", ":4444")
	if err != nil {
		panic(err)
	}
	fmt.Println("Listening on port: 4444")

	err = grpcServer.Serve(listen)
	if err != nil {
		panic(err)
	}

}
