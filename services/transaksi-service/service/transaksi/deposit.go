package account_service

import (
	"context"

	postgres_store "transaksi-service/store/postgres_store/store"
	"transaksi-service/utils/errs"

	"github.com/sirupsen/logrus"
)

type TabungParams struct {
	Nominal       int64  `json:"nominal"`
	NomorRekening string `json:"no_rekening"`
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

	serviceResult.Saldo = storeResult.Akun.Saldo

	return serviceResult, nil
}
