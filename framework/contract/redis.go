package contract

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/gohade/hade/framework"
)

const RedisKey = "hade:redis"

type RedisOption func(container framework.Container, config *RedisConfig) error

type RedisService interface {
	GetClient(option ...RedisOption) (*redis.Client, error)
}

type RedisConfig struct {
	*redis.Options
}

func (config *RedisConfig) UniqKey() string {
	return fmt.Sprintf("%v_%v_%v_%v", config.Addr, config.DB, config.Username, config.Network)
}
