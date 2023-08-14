package main

import (
	"log"

	routeV1 "github.com/ayarushin/go-clean-architecture/api/route/v1"
	"github.com/ayarushin/go-clean-architecture/bootstrap"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := bootstrap.Builder().
		WithEnv().
		WithDB().
		WithTimeout().
		Build()

	server := fiber.New()
	routerApi := server.Group("api")
	routeV1.Setup(app, routerApi)

	log.Fatal(server.Listen(app.Env.ServerAddress))
}
