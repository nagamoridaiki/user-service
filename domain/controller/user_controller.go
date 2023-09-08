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

	res, err := usecase.GetUser(ctx, req)

	if err != nil {
		log.Fatalf("GetUser usecaseの呼び出しエラー: %v", err)
	}

	response := &user.GetUserResponse{
		User: res,
	}

	return response, nil
}

func (s *Server) CreateUser(ctx context.Context, req *user.CreateUserRequest) (*user.CreateUserResponse, error) {
	var res *user.User
	var err error

	// TODO: バリデーションの実装

	res, err = usecase.CreateUser(ctx, req)

	if err != nil {
		log.Fatalf("CreateUser usecaseの呼び出しエラー: %v", err)
	}

	response := &user.CreateUserResponse{
		User: res,
	}

	return response, nil
}
