package user

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
	_ "github.com/go-sql-driver/mysql"
)

type Server struct {
	User UnimplementedUserServiceServer
}

type UserStruct struct {
	UserId       int32
	UserName     string
	UserNameKana *string
	DisplayName  *string
	Email        string
	TwitterId    *string
	LoginId      *string
	Pass         string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (s *Server) GetUser(ctx context.Context, in *GetUserRequest) (*GetUserResponse, error) {
	// ここでユーザー情報を取得する処理を実装する
	// MySQLのDialectオブジェクトを作成

	// データベース接続の設定
	db, err := sql.Open("mysql", "hoge:pass@tcp(127.0.0.1:3307)/member_service?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	dialect := goqu.Dialect("mysql")
	sql, _, err := dialect.From("user").Where(goqu.Ex{"user_id": 2}).ToSQL()
	if err != nil {
		log.Fatal(err)
	}

	// SQLを実行して結果を取得

	// SQLの実行
	rows, err := db.Query(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var user UserStruct
		err := rows.Scan(&user.UserId, &user.UserName, &user.UserNameKana, &user.DisplayName, &user.Email, &user.TwitterId, &user.LoginId, &user.Pass, &user.CreatedAt, &user.UpdatedAt)

		if err != nil {
			panic(err.Error())
		}
		log.Print(user)
		log.Println(&user.UserId, &user.UserName, &user.UserNameKana, &user.DisplayName, &user.Email, &user.TwitterId, &user.LoginId, &user.Pass, &user.CreatedAt, &user.UpdatedAt)

	}

	// 仮のレスポンスを作成して返す
	user := &User{
		UserId:   1,
		UserName: "userName",
		Email:    "email",
		Pass:     "pass",
	}
	response := &GetUserResponse{
		User: user,
	}

	log.Print(response)

	return response, nil
}

func (s *Server) mustEmbedUnimplementedUserServiceServer() {
	panic("必要なエンドポイントはまだ実装されていません")
}
