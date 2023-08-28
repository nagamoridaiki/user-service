package main

import (
	"context"
	"gwitter/user"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	u := user.NewUserServiceClient(conn)

	getUserReq := &user.GetUserRequest{UserId: 1}
	res, err := u.GetUser(context.Background(), getUserReq)
	if err != nil {
		log.Fatalf("Error when calling SyaHello: %s", err)
	}
	log.Printf("Responce from server: %s", res.User)
}
