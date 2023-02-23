package main

import (
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/google/uuid"
	"go-demo-project/pkg/config"
	"go-demo-project/routers"
	"vayuscmdev.ktbcs.co.th/gme/gmxplus/common.git/server"
)

func init() {
	config.Init()
}

func main() {

	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		JSONEncoder:           sonic.Marshal,
		JSONDecoder:           sonic.Unmarshal,
	})
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "*",
		AllowHeaders:     "*",
		AllowCredentials: true}))
	app.Use(recover.New())
	app.Use(requestid.New(requestid.Config{
		Generator:  func() string { return uuid.NewString() },
		ContextKey: "transId",
	}))

	routers.PublicRoutes(app)
	routers.NotFoundRoute(app)

	s := server.NewServer(config.Server().Port)
	s.StartGracefulShutdown(app)
}
