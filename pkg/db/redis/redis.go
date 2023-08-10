package DBRedis

import (
	go_redis "github.com/go-redis/redis"
)

func Init_redis() *go_redis.Client {
	rdb := go_redis.NewClient(&go_redis.Options{
		Addr:     "home.xinxinblog.top:6379",
		Password: "lxx5102G", // no password set
		DB:       0,          // use default DB
	})
	return rdb
}