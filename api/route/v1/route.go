package route

import (
	metricRoute "go-clean-architecture/api/route/v1/metric"
	"go-clean-architecture/bootstrap"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *bootstrap.Application, router fiber.Router) {
	v1 := router.Group("v1")
	metricRoute.New(app, v1)
}
