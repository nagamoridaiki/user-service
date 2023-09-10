package controller

import (
	"context"
	"testing"
	"user-service/user"

	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {

	dummyUser := &user.User{
		UserId:   1,
		UserName: "john_doe",
		Email:    "john@example.com",
		Pass:     "password123",
	}

	controller := &Server{}

	t.Run("GetUser_ValidUser", func(t *testing.T) {

		// ダミーユーザーの情報を使って GetUser を呼び出す
		req := &user.GetUserRequest{UserId: 1}
		res, _ := controller.GetUser(context.Background(), req)

		// 返されたユーザー情報が正しいことを確認
		assert.Equal(t, dummyUser.UserId, res.User.UserId)
		assert.Equal(t, dummyUser.UserName, res.User.UserName)
		assert.Equal(t, dummyUser.Email, res.User.Email)
		assert.Equal(t, dummyUser.Pass, res.User.Pass)
	})
}
