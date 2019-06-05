package main

import (
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_matheDeliveryGRPC "github.com/haffjjj/grpc_learn/server/domain/mathe/delivery/grpc"
	_matheUsecase "github.com/haffjjj/grpc_learn/server/domain/mathe/usercase"
)

func main() {
	//serve grpc(protobuf)
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()

	matheUsecase := _matheUsecase.NewMatheUsecase()

	_matheDeliveryGRPC.NewMatheHandler(srv, matheUsecase)

	reflection.Register(srv)
	fmt.Println("Start server, :4040")
	if e := srv.Serve(listener); e != nil {
		panic(e)
	}
}
