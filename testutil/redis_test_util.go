package golibdataTestUtil

import "github.com/redis/go-redis/v9"

var redisClient *redis.Client

func RedisClient() *redis.Client {
	return redisClient
}
