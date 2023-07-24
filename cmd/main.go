package main

import (
	"log"
	"net"

	"s16-tech-test/config"
	"s16-tech-test/interceptors"
	"s16-tech-test/rpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func CreateGrpcServer() *grpc.Server {
	server := grpc.NewServer(
		grpc.MaxRecvMsgSize(1024*1024*20),
		grpc.MaxSendMsgSize(1024*1024*20),
		grpc.UnaryInterceptor(interceptors.AuthMwUnary()),
	)
	return server
}

func main() {
	// parse config
	cfg, err := config.NewConfig("./config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	// create & register server
	server := CreateGrpcServer()
	rpc.NewOMDBService(server)

	addr := cfg.Server.Host + ":" + cfg.Server.Port
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("could not listen to %s %v", addr, err)
	}
	log.Println("connected to port:", addr)
	reflection.Register(server)

	if err := server.Serve(lis); err != nil {
		log.Fatalf("Server was unable to gracefully shutdown due to err: %+v", err)
	}
}
