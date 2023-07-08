package main

import (
	routeV1 "go-clean-architecture/api/route/v1"
	"go-clean-architecture/bootstrap/app"
	"go-clean-architecture/bootstrap/env"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	env := env.New()
	app := app.NewBuilder().
		WithEnv(env).
		// WithDB().
		Build()

	timeout := time.Duration(env.ContextTimeout) * time.Second

	server := fiber.New()
	routerApi := server.Group("api")
	routeV1.Setup(env, timeout, app.Db, routerApi)

	log.Fatal(server.Listen(env.ServerAddress))
}
