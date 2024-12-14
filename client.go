package main

import (
	"context"
	"log"

	user "github.com/jk.chen/go-rpc-crud/user"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:18089", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := user.NewUserServiceClient(conn)

	user_0 := &user.User{Name: "Tom Hank", Email: "thank@gmail.com"}
	resp, err := client.CreateUser(context.Background(), &user.CreateUserRequest{User: user_0})
	if err != nil {
		log.Fatalf("could not create user: %v", err)
	}
	log.Printf("Created user: %v", resp.User)
}
