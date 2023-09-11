package infra

import (
	"database/sql"
	"fmt"
	"log"

	"user-service/testconfig"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func NewDBConnection() (*sql.DB, error) {
	config := testconfig.NewConfig()

	mysqlUser := config.MySQLUser
	mysqlPassword := config.MySQLPassword
	mysqlHost := config.MySQLHost
	mysqlPort := config.MySQLPort
	mysqlDatabase := config.MySQLDatabase

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", mysqlUser, mysqlPassword, mysqlHost, mysqlPort, mysqlDatabase)

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatalln("データベースの接続エラー: ", err)
		db.Close()
	}

	return db, nil
}
