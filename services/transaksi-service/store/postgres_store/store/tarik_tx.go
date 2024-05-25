package postgres_store

import (
	"context"

	"transaksi-service/store/postgres_store/sqlc"
	"transaksi-service/utils/errs"

	"github.com/sirupsen/logrus"
)

type TarikTxParams struct {
	Nominal       int64  `json:"nominal"`
	NomorRekening string `json:"nomor_rekening"`
}

type TarikTxResult struct {
	Akun      sqlc.DaftarAkun      `json:"akun"`
	Transaksi sqlc.DaftarTransaksi `json:"transaksi"`
}

func (store *PostgresStore) TarikTx(ctx context.Context, arg TarikTxParams) (TarikTxResult, error) {
	const op errs.Op = "postgres_store/TarikTx"

	var result TarikTxResult

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

			return err
		}

		// update saldo
		result.Akun, err = q.UpdateSaldo(ctx, sqlc.UpdateSaldoParams{
			NomorRekening: arg.NomorRekening,
			Saldo:         akun.Saldo - arg.Nominal,
		})
		if err != nil {
			store.logger.WithFields(logrus.Fields{
				"op":    op,
				"scope": "UpdateSaldo",
				"err":   err.Error(),
			}).Error("error!")

			return err
		}

		// Create Transaksi
		result.Transaksi, err = q.CreateTransaksi(ctx, sqlc.CreateTransaksiParams{
			JenisTransaksi: "tarik",
			Nominal:        -arg.Nominal,
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
