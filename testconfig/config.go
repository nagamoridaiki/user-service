package testconfig

import (
	"os"

	"github.com/joho/godotenv"
)

// Config 構造体は設定情報を格納します
type Config struct {
	DatabaseURL       string
	MySQLRootPassword string
	MySQLUser         string
	MySQLPassword     string
	MySQLHost         string
	MySQLPort         string
	MySQLDatabase     string
	TimeZone          string
}

// NewConfig は設定情報を読み込む関数です
func NewConfig() *Config {

	// テスト用
	if err := godotenv.Load(); err != nil {
		return &Config{
			DatabaseURL:       "mysql://hoge:pass@127.0.0.1:3307/member_service?schema=public",
			MySQLRootPassword: "pass",
			MySQLUser:         "hoge",
			MySQLPassword:     "pass",
			MySQLHost:         "127.0.0.1",
			MySQLPort:         "3307",
			MySQLDatabase:     "member_service",
			TimeZone:          "Asia/Tokyo",
		}
	}

	return &Config{
		DatabaseURL:       os.Getenv("DATABASE_URL"),
		MySQLRootPassword: os.Getenv("MYSQL_ROOT_PASSWORD"),
		MySQLUser:         os.Getenv("MYSQL_USER"),
		MySQLPassword:     os.Getenv("MYSQL_PASSWORD"),
		MySQLHost:         os.Getenv("MYSQL_HOST"),
		MySQLPort:         os.Getenv("MYSQL_PORT"),
		MySQLDatabase:     os.Getenv("MYSQL_DATABASE"),
		TimeZone:          os.Getenv("TZ"),
	}
}
