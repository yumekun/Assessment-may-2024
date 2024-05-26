package api_tansaksi

import (
	transaksi_service "transaksi-service/service/transaksi"
)

type Api struct {
	service *transaksi_service.Service
}

func NewApi(transaksiService *transaksi_service.Service) *Api {
	return &Api{
		service: transaksiService,
	}
}
