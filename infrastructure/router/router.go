package router

import (
	"github.com/gin-gonic/gin"
	"go-clean-architecture/infrastructure/controller"
)

func NewRouter(e *gin.Engine, c controller.AppController) *gin.Engine {
	e.Use(gin.Logger())
	e.Use(gin.Recovery())

	e.GET("/metrics", func(context *gin.Context) { c.Metric.GetMetrics(context) })

	return e
}
