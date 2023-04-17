package http

import (
	"fmt"
	"mongoidx/usecases"

	"github.com/gofiber/fiber/v2"
)

type HttpServer struct {
	port    int
	usecase *usecases.Usecase
}

func NewHttpServer(port int, usecase *usecases.Usecase) *HttpServer {
	return &HttpServer{
		port:    port,
		usecase: usecase,
	}
}

func (h *HttpServer) Start() {
	app := fiber.New()

	app.Get("/connections", h.listConnection)
	app.Post("/connections", h.createConnection)
	app.Get("/connections/:id", h.getConnection)
	app.Delete("/connections/:id", h.deleteConnection)

	app.Get("/connect_histories", h.listConnectHistory)

	app.Listen(fmt.Sprintf(":%d", h.port))
}
