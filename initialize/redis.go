package initialize

import (
	"fmt"
	"github.com/hespecial/gin-mall/global"
	"github.com/redis/go-redis/v9"
)

func InitRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", global.Config.Redis.Host, global.Config.Redis.Port),
		Password: global.Config.Redis.Password,
		DB:       global.Config.Redis.Db,
	})
}
