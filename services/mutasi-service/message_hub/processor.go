package request_processor

import (
	"mutasi-service/service"
	postgres_store "mutasi-service/store/postgres_store/store"
	"mutasi-service/store/redis_store"
	"mutasi-service/utils/config"

	"github.com/sirupsen/logrus"
)

type store struct {
	redis *redis_store.RedisStore
}

func newStore(redisStore *redis_store.RedisStore) *store {
	return &store{
		redis: redisStore,
	}
}

type RequestProcessor struct {
	config config.Config
	logger *logrus.Logger

	service *service.Service
	store   *store
}

func NewRequestProcessor(
	config config.Config,
	logger *logrus.Logger,
	postgresStore postgres_store.Store,
	redisStore *redis_store.RedisStore,
) *RequestProcessor {
	service := service.NewService(logger, postgresStore)
	store := newStore(redisStore)

	return &RequestProcessor{
		config: config,
		logger: logger,

		service: service,
		store:   store,
	}
}
