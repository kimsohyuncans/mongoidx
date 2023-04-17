package http

import "github.com/gofiber/fiber/v2"

func (httpServer *HttpServer) listConnectHistory(ctx *fiber.Ctx) error {
	histories, err := httpServer.usecase.ListConnectHistory()
	if err != nil {
		return err
	}

	return ctx.JSON(histories)
}
