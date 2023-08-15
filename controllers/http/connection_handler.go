package http

import (
	"github.com/gofiber/fiber/v2"
)

type createConnectionRequestBody struct {
	Uri string `json:"uri"`
}

func (httpServer *HttpServer) createConnection(ctx *fiber.Ctx) error {
	var requestBody createConnectionRequestBody
	if err := ctx.BodyParser(&requestBody); err != nil {
		return err
	}

	res, err := httpServer.usecase.CreateConnection(requestBody.Uri)
	if err != nil {
		return err
	}

	return ctx.JSON(res)
}

func (httpServer *HttpServer) listConnection(ctx *fiber.Ctx) error {
	res, err := httpServer.usecase.ListConnection()
	if err != nil {
		return err
	}

	return ctx.JSON(res)
}

func (httpServer *HttpServer) getConnection(ctx *fiber.Ctx) error {
	res, err := httpServer.usecase.GetConnection(ctx.Params("id"))
	if err != nil {
		return err
	}

	return ctx.JSON(res)
}

func (httpServer *HttpServer) deleteConnection(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	err := httpServer.usecase.DeleteConnection(id)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusNoContent).JSON(nil)
}
