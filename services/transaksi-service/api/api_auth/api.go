package api_auth

import (
	"transaksi-service/service/auth_service"
)

type Api struct {
	service *auth_service.Service
}

func NewApi(authService *auth_service.Service) *Api {
	return &Api{
		service: authService,
	}
}
