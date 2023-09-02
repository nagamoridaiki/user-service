package user

import (
	"context"
	"log"
	"user-service/infra"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
	_ "github.com/go-sql-driver/mysql"
)

type Server struct {
	User UnimplementedUserServiceServer
}

func (s *Server) GetUser(ctx context.Context, req *GetUserRequest) (*GetUserResponse, error) {

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

	user := User{}
	err = db.QueryRow(sql).Scan(
		&user.UserId,
		&user.UserName,
		&user.UserNameKana,
		&user.DisplayName,
		&user.Email,
		&user.TwitterId,
		&user.LoginId,
		&user.Pass,
	)

	if err != nil {
		log.Fatalln("データベースからの取得失敗エラー: ", err)
	}

	response := &GetUserResponse{
		User: &user,
	}

	log.Print(response)

	return response, nil
}

func (s *Server) mustEmbedUnimplementedUserServiceServer() {
	panic("必要なエンドポイントはまだ実装されていません")
}
