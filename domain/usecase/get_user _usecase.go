package usecase

import (
	"context"
	"user-service/domain/repository"
	"user-service/user"

	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
	_ "github.com/go-sql-driver/mysql"
)

func GetUser(ctx context.Context, req *user.GetUserRequest) (*user.GetUserResponse, error) {

	return repository.GetUser(ctx, req)
}
