package initialize

import (
	"context"
	"fmt"
	"smart-rental/global"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

//var DB *pgxpool.Pool

func InitPostgre() {
	p := global.Config.DB
	dsn :="postgres://%s:%s@%s:%v/%s"
	s := fmt.Sprintf(dsn, p.DBUser, p.DBPassword, p.DBHost, p.DBPort, p.DBName)
	db, err := pgxpool.New(context.Background(), s)
	if err != nil {
		panic(err)
	}
	global.Db = db
}

