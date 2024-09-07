package global

import (
	"smart-rental/pkg/settings"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	Config settings.Config
	Db *pgxpool.Pool
)