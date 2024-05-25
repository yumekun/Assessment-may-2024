package account_service

import (
	postgres_store "transaksi-service/store/postgres_store/store"
	"transaksi-service/utils/config"

	"github.com/sirupsen/logrus"
)

type store struct {
	postgres postgres_store.Store
}

func newStore(postgresStore postgres_store.Store) *store {
	return &store{
		postgres: postgresStore,
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
) *Service {
	store := newStore(postgresStore)

	return &Service{
		config: config,
		logger: logger,
		store:  store,
	}
}
