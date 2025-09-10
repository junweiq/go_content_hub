package main

import (
	"context"
	"database/sql"
	"fmt"
	"go_content_hub/internal/dao"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	ctx := context.Background()
	dsn := "root:handsome@tcp(localhost:3306)/cms_user?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	// 測試連接
	if err := db.Ping(); err != nil {
		log.Fatalf("failed to ping database: %v", err)
	}

	fmt.Println("Connected to database!")

	query := dao.New(db)
	exist, err := query.CheckExist(ctx, sql.NullString{
		String: "admin",
		Valid:  true,
	})
	if err != nil {
		panic(err)
		return
	}
	fmt.Println(exist)
}
