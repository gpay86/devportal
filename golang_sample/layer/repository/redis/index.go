package redis

import (
	"github.com/go-redis/redis/v7"
	"time"
)

func newRedisClient(addr, pass string, db int) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pass,
		DB:       db,
	})
	_, err := client.Ping().Result()
	return client, err
}

func Connection(addr, pass string, db int) *redis.Client {
	redis_conn, err := newRedisClient(addr, pass, db)
	if err != nil {
		panic(err)
	}

	return redis_conn
}

type redisRepositoryImpl struct {
	client *redis.Client
}

func NewRedisRepositoryImpl(client *redis.Client) *redisRepositoryImpl {
	return &redisRepositoryImpl{client}
}

func (r *redisRepositoryImpl) Set(key string, value interface{}, expiration time.Duration) error {
	return r.client.Set(key, value, expiration).Err()
}

func (r *redisRepositoryImpl) Get(key string) string {
	val, err := r.client.Get(key).Result()
	if err != nil {
		return ""
	}
	return val
}

func (r *redisRepositoryImpl) Del(key string) error {
	_, err := r.client.Del(key).Result()
	if err != nil {
		return err
	}
	return nil
}

func (r *redisRepositoryImpl) Exists(key string) bool {
	i, err := r.client.Exists(key).Result()
	if err != nil || i == 0 {
		return false
	}
	return true
}
