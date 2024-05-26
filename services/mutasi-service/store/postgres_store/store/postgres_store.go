package postgres_store

import (
	"database/sql"

	"mutasi-service/store/postgres_store/sqlc"

	"github.com/sirupsen/logrus"
)

type Store interface {
	sqlc.Querier
}

type PostgresStore struct {
	*sqlc.Queries

	logger *logrus.Logger

	db *sql.DB
}

func NewPostgresStore(logger *logrus.Logger, db *sql.DB) Store {
	return &PostgresStore{
		Queries: sqlc.New(db),

		logger: logger,

		db: db,
	}
}
