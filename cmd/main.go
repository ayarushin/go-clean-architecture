package main

import (
	routeV1 "go-clean-architecture/api/route/v1"
	"go-clean-architecture/bootstrap"
	"log"
	"time"

	"github.com/gofiber/fiber"
)

func main() {
	app := bootstrap.App()
	env := app.Env
	db := app.Db

	timeout := time.Duration(env.ContextTimeout) * time.Second

	server := fiber.New()
	routerApi := server.Group("api")
	routeV1.Setup(env, timeout, db, routerApi)

	log.Fatal(server.Listen(env.ServerAddress))
}
