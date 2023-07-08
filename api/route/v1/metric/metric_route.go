package metric

import (
	controller "go-clean-architecture/api/controller/metric"
	"go-clean-architecture/bootstrap/env"
	"go-clean-architecture/domain"
	repository "go-clean-architecture/repository/metric"
	usecases "go-clean-architecture/usecases/metric"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func New(env *env.Env, timeout time.Duration, db *gorm.DB, group fiber.Router) {
	r := repository.New(db, domain.MetricTable)
	c := controller.New(usecases.New(r, timeout), env)
	group.Post("/metric", c.Create)
	group.Get("/metric", c.Fetch)
	group.Get("/metric/:id", c.GetByID)
}
