package controller

import (
	"context"
	"log"
	"user-service/domain/usecase"
	"user-service/user"

	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
	_ "github.com/go-sql-driver/mysql"
)

type Server struct {
	user.UserServiceServer
}

func (s *Server) GetUser(ctx context.Context, req *user.GetUserRequest) (*user.GetUserResponse, error) {

	response, err := usecase.GetUser(ctx, req)
	
	if err != nil {
		log.Fatalf("GetUser usecaseの呼び出しエラー: %v", err)
	}

	return response, nil
}
