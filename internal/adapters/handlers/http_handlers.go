package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/zeabix-cloud-native/workshop-profile-service/internal/adapters/handlers/httpv1"
	"github.com/zeabix-cloud-native/workshop-profile-service/internal/adapters/handlers/httpv2"
	"github.com/zeabix-cloud-native/workshop-profile-service/internal/core/ports"
)

type HTTPHandler struct {
	s   ports.ProfileService
	app *fiber.App
}

func NewHTTPHandler(s ports.ProfileService) *HTTPHandler {
	return &HTTPHandler{
		s:   s,
		app: fiber.New(),
	}
}

func (h *HTTPHandler) Serve(port string) error {
	h.app.Use(cors.New())

	g := h.app.Group("profile-service")

	// V1 apis
	v1 := g.Group("v1")
	handlerV1 := httpv1.NewHttpHandlerV1(h.s, v1)
	handlerV1.Initialize()

	//V2 apis
	v2 := h.app.Group("v2")
	handlerV2 := httpv2.NewHttpHandlerV2(h.s, v2)
	handlerV2.Initialize()

	// Healthcheck
	h.app.Get("/health", h.healthCheck)

	return h.app.Listen(port)
}

func (h *HTTPHandler) healthCheck(c *fiber.Ctx) error {
	c.SendStatus(fiber.StatusOK)
	return nil
}
