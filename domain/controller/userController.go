package controller

import (
	"context"
	"log"
	"user-service/infra"
	"user-service/user"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
	_ "github.com/go-sql-driver/mysql"
)

type Server struct {
	user.UserServiceServer
}

func (s *Server) GetUser(ctx context.Context, req *user.GetUserRequest) (*user.GetUserResponse, error) {

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
		log.Fatalln("データベースからの取得失敗エラー: ", err)
	}

	response := &user.GetUserResponse{
		User: &userData,
	}

	log.Print(response)

	return response, nil
}
