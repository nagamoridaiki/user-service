package repository

import (
	"context"
	"log"
	"user-service/infra"
	"user-service/user"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
	_ "github.com/go-sql-driver/mysql"
)

func GetUser(ctx context.Context, req *user.GetUserRequest) (*user.User, error) {

	db, err := infra.NewDBConnection()
	if err != nil {
		log.Fatalf("データベースの接続エラー: %v", err)
	}
	defer db.Close()

	dialect := goqu.Dialect("mysql")
	sql, _, err := dialect.From("user").Where(goqu.C("user_id").Eq(req.UserId)).ToSQL()
	if err != nil {
		log.Fatal(err)
	}

	var userData user.User
	err = db.QueryRow(sql).Scan(
		&userData.UserId,
		&userData.UserName,
		&userData.UserNameKana,
		&userData.DisplayName,
		&userData.Email,
		&userData.TwitterId,
		&userData.LoginId,
		&userData.Pass,
	)

	if err != nil {
		log.Fatalln("getUser repositoryの取得に失敗しました。: ", err)
	}

	return &userData, nil
}
