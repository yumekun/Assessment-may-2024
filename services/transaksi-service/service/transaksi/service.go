package account_service

import (
	postgres_store "transaksi-service/store/postgres_store/store"
	redis_store "transaksi-service/store/redis_store"
	"transaksi-service/utils/config"

	"github.com/sirupsen/logrus"
)

type store struct {
	postgres postgres_store.Store
	redis    *redis_store.RedisStore
}

func newStore(postgresStore postgres_store.Store, redis_store *redis_store.RedisStore) *store {
	return &store{
		postgres: postgresStore,
		redis:    redis_store,
	}
}

type Service struct {
	config config.Config
	logger *logrus.Logger

	store *store
}

func NewService(
	config config.Config,
	logger *logrus.Logger,
	postgresStore postgres_store.Store,
	redisStore *redis_store.RedisStore,
) *Service {
	store := newStore(postgresStore, redisStore)

	return &Service{
		config: config,
		logger: logger,
		store:  store,
	}
}
