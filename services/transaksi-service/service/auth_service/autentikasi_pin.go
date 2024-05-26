package auth_service

import (
	"context"
	"database/sql"
	"errors"

	"transaksi-service/utils/errs"

	"github.com/sirupsen/logrus"
)

type AutentikasiPinParams struct {
	NomorRekening string `json:"nomor_rekening"`
	Pin           string `json:"pin"`
}

type AutentikasiPinResult struct {
	Authenticated bool `json:"authenticated"`
}

func (service *Service) AutentikasiPin(ctx context.Context, params *AutentikasiPinParams) (*AutentikasiPinResult, error) {
	const op errs.Op = "auth_service/AutentikasiPin"

	serviceResult := &AutentikasiPinResult{}

	service.logger.WithFields(logrus.Fields{
		"op":     op,
		"params": params,
	}).Debug("params!")

	akun, err := service.store.postgres.GetDaftarAkun(ctx, params.NomorRekening)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("nomor rekening tidak dikenali")
		}

		return nil, err
	}

	pelanggan, err := service.store.postgres.GetPelanggan(ctx, akun.IDPelanggan)
	if err != nil {
		return nil, err
	}

	authenticated := true
	if params.Pin != pelanggan.Pin {
		authenticated = false
	}

	serviceResult.Authenticated = authenticated

	return serviceResult, nil
}
