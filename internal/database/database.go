package database

import (
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseConnection() (*gorm.DB)  {
	var err error
	sqlInfo := os.Getenv("CONN_STRING")
	DB, err = gorm.Open(postgres.Open(sqlInfo), &gorm.Config{
		PrepareStmt: false,
	})
	if err != nil {
		panic("failed to connect to database")
	}
	return DB
}






