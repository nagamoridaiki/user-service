# user-service

protoの生成方法
```
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    ./user/user.proto
```

アプリの立ち上げ方
```
docker-compose up -d
go run server.go
```

マイグレーションはsql-migrateを利用
https://github.com/rubenv/sql-migrate
マイグレーションコマンド
```
sql-migrate status
sql-migrate up
sql-migrate down
```

サンプルレコード（Insert文）
```
INSERT INTO user (user_name, user_name_kana, display_name, email, twitter_id, login_id, pass)
VALUES
    ('john_doe', 'じょんどう', 'John Doe', 'john@example.com', 'john_twitter', 'john_login', 'password123'),
    ('jane_smith', 'じぇいんすみす', 'Jane Smith', 'jane@example.com', 'jane_twitter', 'jane_login', 'pass456'),
    ('bob_johnson', 'ぼぶじょんそん', 'Bob Johnson', 'bob@example.com', 'bob_twitter', 'bob_login', 'secret789');
```

動作確認方法
grpcでの起動はBloomRPCを利用
https://github.com/bloomrpc/bloomrpc
