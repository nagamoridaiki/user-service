package controller

import (
	"context"
	"testing"
	"user-service/user"

	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {

	// db, mock, err := sqlmock.New()
	// if err != nil {
	// 	t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	// }
	// defer db.Close()

	dummyUser := &user.User{
		UserId:   1,
		UserName: "john_doe",
		Email:    "john@example.com",
		Pass:     "password123",
	}

	// dbドライバに対する操作のモック定義
	// mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM user WHERE user_id=?`)).
	// 	WithArgs(20).
	// 	WillReturnRows(sqlmock.NewRows([]string{"user_id", "user_name", "email", "pass"}).
	// 		AddRow(dummyUser.UserId, dummyUser.UserName, dummyUser.Email, dummyUser.Pass))

	controller := &Server{}

	t.Run("GetUser_ValidUser", func(t *testing.T) {

		// ダミーユーザーの情報を使って GetUser を呼び出す
		req := &user.GetUserRequest{UserId: 1}
		res, _ := controller.GetUser(context.Background(), req)

		assert.Equal(t, dummyUser.UserId, res.User.UserId)
		assert.Equal(t, dummyUser.UserName, res.User.UserName)
		assert.Equal(t, dummyUser.Email, res.User.Email)
		assert.Equal(t, dummyUser.Pass, res.User.Pass)
	})
}
