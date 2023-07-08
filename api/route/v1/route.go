package route

import (
	metricRoute "go-clean-architecture/api/route/v1/metric"
	"go-clean-architecture/bootstrap/env"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Setup(env *env.Env, timeout time.Duration, db *gorm.DB, router fiber.Router) {
	v1 := router.Group("v1")
	metricRoute.New(env, timeout, db, v1)
}
