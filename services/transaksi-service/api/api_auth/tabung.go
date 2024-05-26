package api_auth

import (
	"context"

	"transaksi-service/service/auth_service"
	transaksi_service "transaksi-service/service/transaksi"

	"github.com/gofiber/fiber/v2"
)

func (api *Api) Tabung(c *fiber.Ctx) error {
	pin := c.Get("Authorization")
	if pin == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(map[string]interface{}{
			"remark": "`Authorization` header is missing",
		})
	}

	var params *transaksi_service.TabungParams
	if err := c.BodyParser(&params); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
			"remark": "failed to parse request body",
		})
	}

	authParams := &auth_service.AutentikasiPinParams{
		NomorRekening: params.NomorRekening,
		Pin:           pin,
	}

	// call service layer
	result, err := api.service.AutentikasiPin(context.Background(), authParams)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{
			"remark": err.Error(),
		})
	}

	if !result.Authenticated {
		return c.Status(fiber.StatusUnauthorized).JSON(map[string]interface{}{
			"remark": "invalid pin",
		})
	}

	return c.Next()
}
