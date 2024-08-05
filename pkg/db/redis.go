package db

import "github.com/qq6049090/TSDaodao-IMSDK/pkg/redis"

func NewRedis(addr string, password string) *redis.Conn {
	return redis.New(addr, password)
}
