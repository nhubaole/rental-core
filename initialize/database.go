package initialize

import (
	"context"
	"smart-rental/internal/dataaccess"

	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

var DB *pgxpool.Pool

func DatabaseConnection() *pgxpool.Pool {
	var err error
	// DB_USER := os.Getenv("DB_USER")
	// DB_PASSWORD := os.Getenv("DB_PASSWORD")
	// DB_HOST := os.Getenv("DB_HOST")
	// DB_PORT := os.Getenv("DB_PORT")
	// DB_NAME := os.Getenv("DB_NAME")
	CONN_STR := os.Getenv("CONN_STRING")
	// connStr := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
	// 	DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)
	DB, err = pgxpool.New(context.Background(), CONN_STR)
	if err != nil {
		panic(err)
	}
	fmt.Println(DB)
	return DB
}

func NewQueries() *dataaccess.Queries {
	return dataaccess.New(DB)
}
