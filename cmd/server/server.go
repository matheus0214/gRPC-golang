package main

import (
	"google.golang.org/grpc"
	"log"
        "net"
        "github.com/matheus0214/grpc/pb"
        "github.com/matheus0214/grpc/services"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

        grpcServer := grpc.NewServer()
        pb.RegisterUserServiceServer(grpcServer, services.NewUserService())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Could not serve: %v", err)
	}
}
