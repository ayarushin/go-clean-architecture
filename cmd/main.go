package main

import (
	"github.com/gin-gonic/gin"
	"go-clean-architecture/infrastructure/datastore"
	"go-clean-architecture/infrastructure/router"
	"go-clean-architecture/registry"
	"log"
)

func main() {
	db := datastore.NewDb()
	r := registry.NewRegistry(db)
	g := gin.Default()
	g = router.NewRouter(g, r.NewAppController())

	if err := g.Run(":3000"); err != nil {
		log.Fatalln(err)
	}
}
