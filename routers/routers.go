package routers

import (
	"fmt"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"go-demo-project/features/sonar"
	"go-demo-project/pkg/config"

	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(app *fiber.App) {
	cfService := config.Service()
	groupV1 := fmt.Sprintf("%s/%s/GoDemo", cfService.Endpoint, "v1")

	routerV1 := app.Group(groupV1)
	helpCheck(routerV1)

	sonar.Router(routerV1)
}

func helpCheck(r fiber.Router) {
	g := r.Group("")
	g.Get("/HelpCheck", func(c *fiber.Ctx) error { return c.SendString("Hello World") })
	g.Get("/monitor", monitor.New(monitor.ConfigDefault))
}
