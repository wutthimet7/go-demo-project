package sonar

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

type (
	handler struct {
		service Service
	}
)

func NewHandler(service Service) *handler {
	return &handler{service}
}

func (h *handler) Sonar(ctx *fiber.Ctx) (err error) {
	if err = h.service.Sonar(ctx); err != nil {
		log.Fatal(err)
		panic(err)
	}

	return ctx.Status(fiber.StatusOK).SendString("Sonar test")
}
