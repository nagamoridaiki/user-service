package controller

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"reflect"
	"regexp"
	"testing"
	"user-service/user"

	"github.com/DATA-DOG/go-sqlmock"
	"google.golang.org/protobuf/proto"
)

var db *sql.DB
var mock sqlmock.Sqlmock

func TestMain(m *testing.M) {
	// テストデータのセットアップ
	var err error
	db, mock, err = sqlmock.New()
	if err != nil {
		fmt.Println("failed to open sqlmock database:", err)
	}
	defer db.Close()

	// モックの設定
	rows := sqlmock.NewRows([]string{
		"user_id", "user_name", "user_name_kana", "display_name", "email",
		"twitter_id", "login_id", "pass",
	}).
		AddRow(10, "test1", "kana1", "display1", "john1@example.com", "twitter1", "login1", "pass1").
		AddRow(20, "test2", "kana2", "display2", "john2@example.com", "twitter2", "login2", "pass2")

	mock.ExpectQuery("SELECT").WillReturnRows(rows)

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `user` (user_id, user_name, user_name_kana, display_name, email, twitter_id, login_id, pass) VALUES (?, ?, ?, ?, ?, ?, ?, ?)")).
		WithArgs(10, "test1", "kana1", "display1", "john1@example.com", "twitter1", "login1", "pass1").
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// テスト実行
	code := m.Run()

	// モックのクエリ実行を検証
	if err := mock.ExpectationsWereMet(); err != nil {
		fmt.Printf("there were unfulfilled expectations: %s", err)
	}

	// テストの結果を返す
	os.Exit(code)
}

func TestUser(t *testing.T) {

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
