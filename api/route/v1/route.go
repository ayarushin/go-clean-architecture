package route

import (
	"go-clean-architecture/bootstrap"
	"time"

	"github.com/gofiber/fiber"
	"gorm.io/gorm"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, routerV1 fiber.Router) {
	publicRouterV1 := routerV1.Group("v1")
	NewMetricRoute(env, timeout, db, publicRouterV1)
}
