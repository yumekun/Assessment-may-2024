package account_service

import (
	"context"

	postgres_store "transaksi-service/store/postgres_store/store"
	"transaksi-service/utils/errs"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type TabungParams struct {
	Nominal       int64  `json:"nominal"`
	NomorRekening string `json:"nomor_rekening"`
}

type TabungResult struct {
	Saldo int64 `json:"saldo"`
}

func (service *Service) Tabung(ctx context.Context, params *TabungParams) (*TabungResult, error) {
	const op errs.Op = "account_service/Tabung"

	serviceResult := &TabungResult{}

	service.logger.WithFields(logrus.Fields{
		"op":     op,
		"params": params,
	}).Debug("params!")

	storeResult, err := service.store.postgres.TabungTx(ctx, postgres_store.TabungTxParams{
		Nominal:       params.Nominal,
		NomorRekening: params.NomorRekening,
	})
	if err != nil {
		return nil, err
	}
	// ID             string `json:"id"`
	// NomorRekening  string `json:"nomor_rekening"`
	// JenisTransaksi string `json:"jenis_transaksi"`
	// Nominal        int64  `json:"nominal"`
	err = service.store.redis.AddToStream(ctx, service.config.RedisMutasiRequestStream, map[string]interface{}{
		"id":              uuid.NewString(),
		"nomor_rekening":  params.NomorRekening,
		"jenis_transaksi": storeResult.Transaksi.JenisTransaksi,
		"nominal":         params.Nominal,
	})
	if err != nil {
		return nil, err
	}

	serviceResult.Saldo = storeResult.Akun.Saldo

	return serviceResult, nil
}
