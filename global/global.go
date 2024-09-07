package global

import (
	"smart-rental/pkg/settings"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

var (
	Config settings.Config
	Db *pgxpool.Pool
	S3 *s3.Client
	Redis *redis.Client
)