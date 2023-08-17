package DBRedis

import (
	"fmt"
	"strconv"

	"github.com/go-redis/redis"
)

type RedisDB struct {
	RDB *redis.Client
}

func (r *RedisDB) InitRedisDB() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "192.168.45.128:6379",
		Password: "lxx123",
		DB:       0,
	})
	r.RDB = rdb
}

// 设置client在线状态
func (r *RedisDB) SetClientStatus(client_id uint64, status bool) {
	var key string
	var value bool
	if status == true {
		temp := strconv.FormatUint(client_id, 10)
		key = fmt.Sprintf("%s$status", temp)
		value = true
	} else if status == false {
		temp := strconv.FormatUint(client_id, 10)
		key = fmt.Sprintf("%s$status", temp)
		value = false
	}
	err := r.RDB.Set(key, value, 0).Err()
	if err != nil {
		fmt.Println("err 2133:", err)
	}
}

func (r *RedisDB) GetClientStauts(client_id uint64) bool {
	temp := strconv.FormatUint(client_id, 10)
	key := fmt.Sprintf("%s$status", temp)
	value, err := r.RDB.Get(key).Result()
	if err != nil {
		fmt.Println("email_code reserve fail")
	}
	if len(value) == 0 {
		return false
	}
	return true
}
