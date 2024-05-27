package account_service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	postgres_store "transaksi-service/store/postgres_store/store"
	"transaksi-service/utils/errs"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type TarikParams struct {
	Nominal       int64  `json:"nominal"`
	NomorRekening string `json:"nomor_rekening"`
}

type TarikResult struct {
	Saldo int64 `json:"saldo"`
}

func (service *Service) Tarik(ctx context.Context, params *TarikParams) (*TarikResult, error) {
	const op errs.Op = "account_service/Tarik"

	serviceResult := &TarikResult{}

	service.logger.WithFields(logrus.Fields{
		"op":     op,
		"params": params,
	}).Info("params!")

	account, err := service.store.postgres.GetDaftarAkun(ctx, params.NomorRekening)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("nomor rekening tidak dikenali")
		}

		return nil, err
	}

	if account.Saldo < params.Nominal {
		return nil, fmt.Errorf("saldo tidak cukup")
	}

	result, err := service.store.postgres.TarikTx(ctx, postgres_store.TarikTxParams{
		Nominal:       params.Nominal,
		NomorRekening: params.NomorRekening,
	})
	if err != nil {
		return nil, err
	}
	err = service.store.redis.AddToStream(ctx, service.config.RedisMutasiRequestStream, map[string]interface{}{
		"id":              uuid.NewString(),
		"nomor_rekening":  params.NomorRekening,
		"jenis_transaksi": result.Transaksi.JenisTransaksi,
		"nominal":         -params.Nominal,
	})
	if err != nil {
		return nil, err
	}

	serviceResult.Saldo = result.Akun.Saldo

	return serviceResult, nil
}
