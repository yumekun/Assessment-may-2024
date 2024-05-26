package service

import (
	postgres_store "mutasi-service/store/postgres_store/store"

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
	logger *logrus.Logger

	store *store
}

func NewService(
	logger *logrus.Logger,
	postgresStore postgres_store.Store,
) *Service {
	store := newStore(postgresStore)

	return &Service{
		logger: logger,

		store: store,
	}
}
