package main

import (
	routeV1 "go-clean-architecture/api/route/v1"
	"go-clean-architecture/bootstrap"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	app := bootstrap.App()
	env := app.Env
	db := app.Db

	timeout := time.Duration(env.ContextTimeout) * time.Second

	gin := gin.Default()
	routerV1 := gin.Group("v1") // migrate to fasthttp
	routeV1.Setup(env, timeout, db, routerV1)

	gin.Run(env.ServerAddress)

	if err := gin.Run(env.ServerAddress); err != nil {
		log.Fatalln(err)
	}
}
