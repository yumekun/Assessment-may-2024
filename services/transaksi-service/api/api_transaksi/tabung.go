package api_tansaksi

import (
	"context"

	transaksi "transaksi-service/service/transaksi"

	"github.com/gofiber/fiber/v2"
)

func (api *Api) Tabung(c *fiber.Ctx) error {
	var params *transaksi.TabungParams
	if err := c.BodyParser(&params); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
			"remark": "failed to parse request body",
		})
	}

	// call service layer
	result, err := api.service.Tabung(context.Background(), params)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{
			"remark": err.Error(),
		})
	}

	// tidy up response
	response := map[string]interface{}{
		"saldo": result.Saldo,
	}

	return c.JSON(response)
}
