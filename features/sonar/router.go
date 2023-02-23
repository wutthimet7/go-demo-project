package sonar

import (
	"github.com/gofiber/fiber/v2"
)

func Router(r fiber.Router) {
	service := NewService()
	handler := NewHandler(service)

	r.Get("/Sonar", handler.Sonar)
}
