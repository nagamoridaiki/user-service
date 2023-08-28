package user

import (
	"context"
	"log"
)

type Server struct {
	User UnimplementedUserServiceServer
}

func (s *Server) GetUser(ctx context.Context, in *GetUserRequest) (*GetUserResponse, error) {
	// ここでユーザー情報を取得する処理を実装する
	// 例えば、データベースからユーザー情報を取得するなど

	// 仮のレスポンスを作成して返す
	user := &User{
		UserId:   1,
		UserName: "example_user",
		Email:    "user@example.com",
		Pass:     "test",
	}
	response := &GetUserResponse{
		User: user,
	}

	log.Printf("Receive userId from client: %v", in.UserId)

	return response, nil
}

func (s *Server) mustEmbedUnimplementedUserServiceServer() {
	panic("必要なエンドポイントはまだ実装されていません")
}
