package route

import (
	"go-clean-architecture/api/controller"
	"go-clean-architecture/bootstrap"
	"go-clean-architecture/domain"
	"go-clean-architecture/repository"
	"go-clean-architecture/usecases"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func NewMetricRoute(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, group fiber.Router) {
	mr := repository.NewMetricRepository(db, domain.MetricTable)
	mc := &controller.MetricController{
		MetricUsecase: usecases.NewMetricUsecase(mr, timeout),
		Env:           env,
	}
	group.Post("/metric", mc.Create)
	group.Get("/metric", mc.Fetch)
	group.Get("/metric/:id", mc.GetByID)
}
