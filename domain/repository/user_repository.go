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

func GetUserByName(ctx context.Context, UserName string) (*user.User, error) {

	db, err := infra.NewDBConnection()
	if err != nil {
		log.Fatalf("データベースの接続エラー: %v", err)
	}
	defer db.Close()

	dialect := goqu.Dialect("mysql")
	sql, _, err := dialect.From("user").Where(goqu.C("user_name").Eq(UserName)).ToSQL()
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
		log.Fatalln("getUserByNameの取得に失敗しました。: ", err)
	}

	return &userData, nil
}

func CreateUser(ctx context.Context, req *user.CreateUserRequest) {

	db, err := infra.NewDBConnection()
	if err != nil {
		log.Fatalf("データベースの接続エラー: %v", err)
	}
	defer db.Close()

	dialect := goqu.Dialect("mysql")

	insertSQL := dialect.Insert("user").Rows(
		goqu.Record{
			"user_name":      req.UserName,
			"user_name_kana": req.UserNameKana,
			"display_name":   req.DisplayName,
			"email":          req.Email,
			"twitter_id":     req.TwitterId,
			"login_id":       req.LoginId,
			"pass":           req.Pass,
		},
	)
	sqlString, args, err := insertSQL.ToSQL()

	_, err = db.Exec(sqlString, args...)

	if err != nil {
		log.Fatalln("CreateUserの作成に失敗しました。: ", err)
	}
}
