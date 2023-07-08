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
	mr := repository.New(db, domain.MetricTable)
	mc := &controller.MetricController{
		MetricUsecase: usecases.New(mr, timeout),
		Env:           env,
	}
	group.Post("/metric", mc.Create)
	group.Get("/metric", mc.Fetch)
	group.Get("/metric/:id", mc.GetByID)
}
