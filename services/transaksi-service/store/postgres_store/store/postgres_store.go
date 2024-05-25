package postgres_store

import (
	"context"
	"database/sql"

	"transaksi-service/store/postgres_store/sqlc"
	"transaksi-service/utils/errs"

	"github.com/sirupsen/logrus"
)

type Store interface {
	sqlc.Querier
	TabungTx(ctx context.Context, arg TabungTxParams) (TabungTxResult, error)
	TarikTx(ctx context.Context, arg TarikTxParams) (TarikTxResult, error)
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

func (store *PostgresStore) execTx(ctx context.Context, fn func(*sqlc.Queries) error) error {
	const op errs.Op = "postgres_store/execTx"

	// ensures that only one goroutine can start a transaction at a time,
	// preventing potential race conditions during the setup of the transaction.

	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		store.logger.WithFields(logrus.Fields{
			"op":    op,
			"scope": "BeginTx",
			"err":   err.Error(),
		}).Error("failed to begin tx!")

		return err
	}

	q := sqlc.New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			store.logger.WithFields(logrus.Fields{
				"op":    op,
				"scope": "Rollback",
				"err":   err.Error(),
			}).Error("failed to rollback tx!")

			return err
		}
		return err
	}

	return tx.Commit()
}
