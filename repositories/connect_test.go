package repositories

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func ConnectDB() (db *sqlx.DB, err error) {

	dbUser := os.Getenv("MYSQL_USER")
	dbPassword := os.Getenv("MYSQL_PASSWORD")
	dbHost := "127.0.0.1"
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("MYSQL_DATABASE")
	// DBの接続
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err = sqlx.Connect("mysql", dsn)

	if err != nil {
		log.Fatal(err)
	}

	return
}
