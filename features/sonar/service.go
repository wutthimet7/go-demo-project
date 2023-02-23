package sonar

import (
	"github.com/gofiber/fiber/v2"
)

type (
	Service interface {
		Sonar(ctx *fiber.Ctx) error
	}

	service struct {
	}
)

func NewService() Service {
	return &service{}
}

func (s *service) Sonar(ctx *fiber.Ctx) error {
	return nil
}
