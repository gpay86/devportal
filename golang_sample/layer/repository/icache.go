package repository

import (
	"time"
)

// IRedis -
type IRedis interface {
	Set(key string, value interface{}, expiration time.Duration) error
	Get(key string) string
	Del(key string) error
	Exists(key string) bool
}
