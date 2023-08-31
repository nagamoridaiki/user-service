package user

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
	_ "github.com/go-sql-driver/mysql"
)

type Server struct {
	User UnimplementedUserServiceServer
}

func (s *Server) GetUser(ctx context.Context, in *GetUserRequest) (*GetUserResponse, error) {
	// ここでユーザー情報を取得する処理を実装する
	// MySQLのDialectオブジェクトを作成

	// データベース接続の設定
	db, err := sql.Open("mysql", "hoge:pass@tcp(127.0.0.1:3307)/member_service?parseTime=true")
	if err != nil {
		panic(err.Error())
	}

	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("データベース接続完了")

	// SQLを実行して結果を取得

	// SQLの実行
	user := User{}
	err = db.QueryRow(`select * from user where user_id=1`).Scan(
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
