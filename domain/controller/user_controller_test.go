package controller

import (
	"context"
	"database/sql"
	"reflect"
	"testing"
	"user-service/user"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"google.golang.org/protobuf/proto"
)

var db *sql.DB
var mock sqlmock.Sqlmock

func setupTestDatabase() (*sqlx.DB, error) {

	db, err := sqlx.Open("mysql", "hoge:pass@tcp(127.0.0.1:3307)/member_service")
	if err != nil {
		return nil, err
	}

	// テーブルの作成
	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS user (
            user_id INT(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT'ユーザーID',
			user_name VARCHAR(254) UNIQUE NOT NULL COMMENT'ユーザー名',
			user_name_kana VARCHAR(254) COMMENT'ユーザー名（かな）',
			display_name VARCHAR(255) COMMENT '画面上表示名',
			email VARCHAR(255) NOT NULL COMMENT'メールアドレス',
			twitter_id VARCHAR(100) UNIQUE COMMENT 'TwietterのID',
			login_id VARCHAR(254) UNIQUE COMMENT 'ログインID',
			pass VARCHAR(254) UNIQUE NOT NULL COMMENT 'ユーザー名',

			PRIMARY KEY(user_id)
        )
    `)
	if err != nil {
		return nil, err
	}

	// ダミーレコードの挿入
	_, err = db.Exec(`INSERT INTO user
    (user_id, user_name, user_name_kana, display_name, email, twitter_id, login_id, pass)
    VALUES (10, 'test1', 'kana1', 'display1', 'john1@example.com', 'twitter1', 'login1', 'pass1')`)

	if err != nil {
		return nil, err
	}

	return db, nil
}

func TestUser(t *testing.T) {

	db, err := setupTestDatabase()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	// テスト対象のコントローラーを作成
	controller := &Server{} // あなたの実際のコントローラーに置き換える必要があります

	// テスト用のリクエストを作成
	req := &user.GetUserRequest{UserId: 10}

	// GetUserメソッドを呼び出し、結果を取得
	response, err := controller.GetUser(context.Background(), req)
	if err != nil {
		t.Fatalf("GetUser failed: %v", err)
	}

	// 期待される結果と比較
	expectedUser := &user.User{
		UserId:       10,
		UserName:     "test1",
		UserNameKana: proto.String("kana1"),    // ポインタ型の文字列ポインタを使用
		DisplayName:  proto.String("display1"), // ポインタ型の文字列ポインタを使用
		Email:        "john1@example.com",
		TwitterId:    proto.String("twitter1"), // ポインタ型の文字列ポインタを使用
		LoginId:      proto.String("login1"),   // ポインタ型の文字列ポインタを使用
		Pass:         "pass1",
	}

	if !reflect.DeepEqual(response.User, expectedUser) {
		t.Errorf("Unexpected result. Got: %+v, Expected: %+v", response.User, expectedUser)
	}
}
