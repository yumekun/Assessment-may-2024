package service

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"

	"mutasi-service/store/postgres_store/sqlc"
	// "mutasi-service/utils/cast"
	"mutasi-service/utils/errs"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type CreateMutasiParams struct {
	ID             string `json:"id"`
	JenisTransaksi string `json:"jenis_transaksi"`
	NomorRekening  string `json:"nomor_rekening"`
	Nominal        int64  `json:"nominal"`
}
type Mutasi struct {
	ID             string `json:"id"`
	JenisTransaksi string `json:"jenis_transaksi"`
	NomorRekening  string `json:"nomor_rekening"`
	Nominal        int64  `json:"nominal"`
}

type CreateMutasiResult struct {
	Mutasi Mutasi `json:"saldo"`
}

func (service *Service) CreateMutasi(ctx context.Context, params *CreateMutasiParams) (*CreateMutasiResult, error) {
	const op errs.Op = "service/CreateMutasi"

	serviceResult := &CreateMutasiResult{}

	service.logger.WithFields(logrus.Fields{
		"op":     op,
		"params": params,
	}).Debug("params!")
	// ID             string `json:"id"`
	// JenisTransaksi string `json:"jenis_transaksi"`
	// Nominal        int64  `json:"nominal"`
	// NomorRekening  string `json:"nomor_rekening"`
	mutasi, err := service.store.postgres.CreateMutasi(ctx, sqlc.CreateMutasiParams{
		ID:             uuid.NewString(),
		NomorRekening:  sql.NullString{Valid: true, String: params.NomorRekening},
		JenisTransaksi: sql.NullString{Valid: true, String: params.JenisTransaksi},
		Nominal:        sql.NullInt64{Valid: true, Int64: params.Nominal},
	})
	if err != nil {
		return nil, err
	}

	serviceResult.Mutasi = func(mutasi sqlc.Mutasi) Mutasi {
		mappedMutasi := Mutasi{
			ID:             mutasi.ID,
			JenisTransaksi: mutasi.JenisTransaksi.String,
			NomorRekening:  mutasi.NomorRekening.String,
			Nominal:        mutasi.Nominal.Int64,
		}

		return mappedMutasi
	}(mutasi)

	return serviceResult, nil
}

func (service *Service) NewCreateMutasiParamsFromMap(mapParams map[string]interface{}) (*CreateMutasiParams, error) {
	const op errs.Op = "service/NewCreateMutasiParamsFromMap"

	params, err := func(mapParams map[string]interface{}) (*CreateMutasiParams, error) {
		params := &CreateMutasiParams{}
		id, ok := mapParams["id"].(string)
		if !ok {
			return nil, fmt.Errorf("`id` as string, got type %T", mapParams["id"])
		}
		params.ID = id

		jenisTransaksi, ok := mapParams["jenis_transaksi"].(string)
		if !ok {
			return nil, fmt.Errorf("`jenis_transaksi` as string, got type %T", mapParams["jenis_transaksi"])
		}
		params.JenisTransaksi = jenisTransaksi

		nomorRekening, ok := mapParams["nomor_rekening"].(string)
		if !ok {
			return nil, fmt.Errorf("`nomor_rekening` as string, got type %T", mapParams["nomor_rekening"])
		}
		params.NomorRekening = nomorRekening

		nominalString, ok := mapParams["nominal"].(string)
		if !ok {
			return nil, fmt.Errorf("`nominal` as string, got type %T", mapParams["nominal"])
		}
		nominal, err := strconv.ParseInt(nominalString, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to cast `nominal` as int64: %s", err.Error())
		}
		params.Nominal = int64(nominal)

		return params, nil
	}(mapParams)
	if err != nil {
		e := fmt.Errorf("failed to extract value: %s", err.Error())

		service.logger.WithFields(logrus.Fields{
			"op":    op,
			"scope": "Data Extraction",
			"err":   e.Error(),
		}).Error("error!")

		return nil, e
	}

	return params, nil
}
