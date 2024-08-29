package cache

import "time"

// Cache 缓存接口
type Cache interface {
	// Set 设置key value
	Set(key string, value string) error
	// 删除key
	Delete(key string) error
	// SetAndExpire  设置key value 并支持过期时间
	SetAndExpire(key string, value string, expire time.Duration) error
	// 获取key对应的值
	Get(key string) (string, error)

	// LPUSH 向List添加元素
	LPUSH(key string, values ...interface{}) (int64, error)

	// LPop  获取List的第一个元素
	LPop(key string) (string, error)

	// SAdd 向Set添加元素
	SAdd(key string, members ...interface{}) error

	// SMembers 获取Set中的所有元素
	SMembers(key string) ([]string, error)
}
