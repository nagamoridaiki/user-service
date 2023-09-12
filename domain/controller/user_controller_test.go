package controller

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"testing"
	"user-service/testconfig"
	"user-service/user"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"google.golang.org/protobuf/proto"
)

var db *sql.DB
var mock sqlmock.Sqlmock

func dummyTestUser() *user.User {
	return &user.User{
		UserId:       10,
		UserName:     "test1",
		UserNameKana: proto.String("kana1"),
		DisplayName:  proto.String("display1"),
		Email:        "john1@example.com",
		TwitterId:    proto.String("twitter1"),
		LoginId:      proto.String("login1"),
		Pass:         "pass1",
	}
}

func setupTestDatabase() (*sqlx.DB, error) {

	conf := testconfig.NewConfig()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", conf.MySQLUser, conf.MySQLPassword, conf.MySQLHost, conf.MySQLPort, conf.MySQLDatabase)

	db, err := sqlx.Open("mysql", dsn)
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

	du := dummyTestUser()
	testValue := fmt.Sprintf("%d, '%s', '%s', '%s', '%s', '%s', '%s', '%s'",
		du.UserId, du.UserName, *du.UserNameKana, *du.DisplayName, du.Email, *du.TwitterId, *du.LoginId, du.Pass)

	_, err = db.Exec(`INSERT INTO user
    (user_id, user_name, user_name_kana, display_name, email, twitter_id, login_id, pass)
    VALUES (` + testValue + `)`)

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

	if !reflect.DeepEqual(response.User, dummyTestUser()) {
		t.Errorf("Unexpected result. Got: %+v, Expected: %+v", response.User, dummyTestUser())
	}

	// テストデータのクリーンアップ
	_, err = db.Exec("DELETE FROM user WHERE user_id = ?", req.UserId)
	if err != nil {
		t.Fatalf("Failed to clean up test data: %v", err)
	}
}
