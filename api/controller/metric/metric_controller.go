package metric

import (
	"go-clean-architecture/bootstrap/env"
	"go-clean-architecture/domain"
	"go-clean-architecture/domain/responses"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type MetricController struct {
	MetricUsecase domain.MetricUsecase
	Env           *env.Env
}

func (mc *MetricController) Create(c *fiber.Ctx) error {
	var metric domain.Metric

	err := c.BodyParser(&metric)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.ErrorResponse{Message: err.Error()})
	}

	err = mc.MetricUsecase.Create(c.Context(), &metric)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.ErrorResponse{Message: err.Error()})
	}

	return c.Status(http.StatusOK).JSON(responses.SuccessResponse{
		Message: "Task created successfully",
	})
}

func (mc *MetricController) Fetch(c *fiber.Ctx) error {
	metrics, err := mc.MetricUsecase.Fetch(c.Context())
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.ErrorResponse{Message: err.Error()})
	}

	return c.Status(http.StatusOK).JSON(metrics)
}

func (mc *MetricController) GetByID(c *fiber.Ctx) error {
	id := c.Params(":id")

	metric, err := mc.MetricUsecase.GetByID(c.Context(), id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.ErrorResponse{Message: err.Error()})
	}

	return c.Status(http.StatusOK).JSON(metric)
}
