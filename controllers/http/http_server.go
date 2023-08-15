package http

import (
	"fmt"
	"mongoidx/usecases"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type HttpServer struct {
	app     *fiber.App
	port    int
	usecase *usecases.Usecase
}

func NewHttpServer(port int, usecase *usecases.Usecase) *HttpServer {
	httpServer := &HttpServer{
		port:    port,
		usecase: usecase,
		app:     fiber.New(),
	}

	httpServer.registerRoutes(httpServer.app)
	return httpServer
}

func (h *HttpServer) Start() {
	h.app.Listen(fmt.Sprintf(":%d", h.port))
}

func (h *HttpServer) Test(req *http.Request) (*http.Response, error) {
	return h.app.Test(req)
}

func (h *HttpServer) registerRoutes(app *fiber.App) {
	app.Get("/connections", h.listConnection)
	app.Post("/connections", h.createConnection)
	app.Get("/connections/:id", h.getConnection)
	app.Delete("/connections/:id", h.deleteConnection)

	app.Get("/connect_histories", h.listConnectHistory)
}
