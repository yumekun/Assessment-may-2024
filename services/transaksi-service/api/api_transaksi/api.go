package api_tansaksi

import (
	account_service "transaksi-service/service/transaksi"
)

type Api struct {
	service *account_service.Service
}

func NewApi(accountService *account_service.Service) *Api {
	return &Api{
		service: accountService,
	}
}
