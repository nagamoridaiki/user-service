package main

import (
	"gwitter/user"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := user.Server{}
	grpcServer := grpc.NewServer()
	user.RegisterUserServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

func (s *Server) mustEmbedUnimplementedUserServiceServer() {
	panic("必要なエンドポイントはまだ実装されていません")
}
