package usecase

import (
	"context"
	"log"
	"user-service/domain/repository"
	"user-service/user"

	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
	_ "github.com/go-sql-driver/mysql"
)

func GetUser(ctx context.Context, req *user.GetUserRequest) (*user.User, error) {

	user, err := repository.GetUser(ctx, req)

	if err != nil {
		log.Fatalf("GetUser repositoryの呼び出しエラー: %v", err)
	}

	return user, nil
}
