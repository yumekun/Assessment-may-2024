package postgres_store

import (
	"context"
	"database/sql"
	"errors"

	"transaksi-service/store/postgres_store/sqlc"
	"transaksi-service/utils/errs"

	"github.com/sirupsen/logrus"
)

type TabungTxParams struct {
	Nominal       int64  `json:"nominal"`
	NomorRekening string `json:"nomor_rekening"`
}

type TabungTxResult struct {
	Akun      sqlc.DaftarAkun      `json:"akun"`
	Transaksi sqlc.DaftarTransaksi `json:"transaksi"`
}

func (store *PostgresStore) TabungTx(ctx context.Context, arg TabungTxParams) (TabungTxResult, error) {
	const op errs.Op = "postgres_store/TabungTx"

	var result TabungTxResult

	err := store.execTx(ctx, func(q *sqlc.Queries) error {
		var err error

		// get akun
		akun, err := q.GetDaftarAkun(ctx, arg.NomorRekening)
		if err != nil {
			store.logger.WithFields(logrus.Fields{
				"op":    op,
				"scope": "GetDaftarAkun",
				"err":   err.Error(),
			}).Error("error!")

			if err == sql.ErrNoRows {
				return errors.New("nomor rekening tidak dikenali")
			}

			return err
		}

		// update saldo
		result.Akun, err = q.UpdateSaldo(ctx, sqlc.UpdateSaldoParams{
			NomorRekening: arg.NomorRekening,
			Saldo:         akun.Saldo + arg.Nominal,
		})
		if err != nil {
			store.logger.WithFields(logrus.Fields{
				"op":    op,
				"scope": "UpdateSaldo",
				"err":   err.Error(),
			}).Error("error!")

			return err
		}

		// create transaksi
		result.Transaksi, err = q.CreateTransaksi(ctx, sqlc.CreateTransaksiParams{
			JenisTransaksi: "tabung",
			Nominal:        arg.Nominal,
			NomorRekening:  arg.NomorRekening,
		})
		if err != nil {
			store.logger.WithFields(logrus.Fields{
				"op":    op,
				"scope": "CreateTransaksi",
				"err":   err.Error(),
			}).Error("error!")

			return err
		}

		return err
	})

	return result, err
}
