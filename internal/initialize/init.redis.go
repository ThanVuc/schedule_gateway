package initialize

import (
	"fmt"
	"schedule_gateway/global"
	"strconv"

	"github.com/thanvuc/go-core-lib/cache"
)

func InitRedis() {
	redisConfig := global.Config.Redis
	println("HOST: " + fmt.Sprintf("%s:%s", redisConfig.Host, strconv.Itoa(redisConfig.Port)))
	redisClient := cache.NewRedisCache(cache.Config{
		Addr:     fmt.Sprintf("%s:%s", redisConfig.Host, strconv.Itoa(redisConfig.Port)),
		DB:       redisConfig.DB,
		Password: redisConfig.Password,
		PoolSize: redisConfig.PoolSize,
		MinIdle:  redisConfig.MinIdle,
	})

	if err := redisClient.Ping(); err != nil {
		global.Logger.Error("Failed to connect to Redis", "")
		panic(fmt.Sprintf("Failed to connect to Redis: %v", err))
	} else {
		global.Logger.Info("Redis connection established successfully", "")
	}

	global.RedisDb = redisClient
}
