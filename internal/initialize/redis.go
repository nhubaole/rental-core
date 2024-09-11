package initialize

import (
	"smart-rental/global"

	"github.com/redis/go-redis/v9"
)

func InitRedis() {
	client := redis.NewClient(&redis.Options{
		Addr:     global.Config.DB.DBHost,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	global.Redis = client
}