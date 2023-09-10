package controller

import (
	"context"
	"log"
	"user-service/domain/controller/form"
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
	var user_form error

	user_form = form.UserInputForm(req)
	if user_form != nil {
		log.Fatalf("Userの入力に誤りがあります: %v", user_form)
		return nil, err
	}

	res, err = usecase.CreateUser(ctx, req)

	if err != nil {
		log.Fatalf("CreateUser usecaseの呼び出しエラー: %v", err)
	}

	response := &user.CreateUserResponse{
		User: res,
	}

	return response, nil
}
