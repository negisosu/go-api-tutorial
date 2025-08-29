package main // mainであることの宣言

import (
	"fmt"
	"log"
	"os"
	"time"

	"go-todo-app/controllers"
	"go-todo-app/router"
	"go-todo-app/services"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

// 受け取ったデータからこっちで使用する構造体
type Todo struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func main() {
	// データベース
	_ = godotenv.Load()

	dbUser := os.Getenv("MYSQL_USER")
	dbPassword := os.Getenv("MYSQL_PASSWORD")
	dbHost := "127.0.0.1"
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("MYSQL_DATABASE")
	// DBの接続
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := sqlx.Connect("mysql", dsn)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// サービス
	service := services.NewTodoService(db)

	// コントローラー
	controller := controllers.NewTodoController(service)

	fmt.Println("Successfully connected to the database")

	//　ルーター
	e := router.NewRouter(controller)
	// サーバーの起動
	log.Println("server start at port 8080")

	if err := e.Start(":8080"); err != nil {
		log.Fatal(err)
	}
}
