package main

import (
	routeV1 "go-clean-architecture/api/route/v1"
	bootstrap "go-clean-architecture/bootstrap/app"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := bootstrap.New()
	env := app.Env
	db := app.Db

	timeout := time.Duration(env.ContextTimeout) * time.Second

	server := fiber.New()
	routerApi := server.Group("api")
	routeV1.Setup(env, timeout, db, routerApi)

	log.Fatal(server.Listen(env.ServerAddress))
}
