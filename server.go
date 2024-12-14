package main

import (
	"context"
	"log"
	"net"

	user "github.com/jk.chen/go-rpc-crud/user"
	"google.golang.org/grpc"
)

type userServiceServer struct {
	user.UnimplementedUserServiceServer
}

func (s *userServiceServer) CreateUser(ctx context.Context, req *user.CreateUserRequest) (*user.CreateUserResponse, error) {
	return &user.CreateUserResponse{User: req.User}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":18089")

	if err != nil {
		log.Fatalf("canoot create listener %s", err)
	}

	grpcServer := grpc.NewServer()
	user.RegisterUserServiceServer(grpcServer, &userServiceServer{})

	log.Println("gRPC server listening on :18089")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("unable to register invoice server %s", err)
	}
}
