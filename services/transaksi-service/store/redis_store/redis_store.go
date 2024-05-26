package redis_store

import (
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

type RedisStore struct {
	logger *logrus.Logger

	client *redis.Client
}

func NewRedisStore(logger *logrus.Logger, redisClient *redis.Client) *RedisStore {
	return &RedisStore{
		logger: logger,

		client: redisClient,
	}
}
