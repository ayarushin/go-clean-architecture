package route

import (
	metricRoute "go-clean-architecture/api/route/v1/metric"
	"go-clean-architecture/bootstrap/env"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Setup(env *env.Env, timeout time.Duration, db *gorm.DB, routerV1 fiber.Router) {
	publicRouterV1 := routerV1.Group("v1")
	metricRoute.New(env, timeout, db, publicRouterV1)
}
