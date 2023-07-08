package metric

import (
	controller "go-clean-architecture/api/controller/metric"
	"go-clean-architecture/bootstrap"
	"go-clean-architecture/domain"
	repository "go-clean-architecture/repository/metric"
	usecases "go-clean-architecture/usecases/metric"

	"github.com/gofiber/fiber/v2"
)

func New(app *bootstrap.Application, group fiber.Router) {
	r := repository.New(app.Db, domain.MetricTable)
	u := usecases.New(r, app.Timeout)
	c := controller.New(u)
	group.Post("/metric", c.Create)
	group.Get("/metric", c.Fetch)
	group.Get("/metric/:id", c.GetByID)
}
