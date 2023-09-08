package usecase

import (
	"context"
	"log"
	"user-service/domain/repository"
	"user-service/user"

	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
	_ "github.com/go-sql-driver/mysql"
)

func CreateUser(ctx context.Context, req *user.CreateUserRequest) (*user.User, error) {

	repository.CreateUser(ctx, req)

	res, err := repository.GetUserByName(ctx, req.UserName)
	if err != nil {
		log.Fatalf("GetUser repositoryの呼び出しエラー: %v", err)
	}

	return res, err
}
