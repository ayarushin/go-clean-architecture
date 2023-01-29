package controller

import (
	"go-clean-architecture/bootstrap"
	"go-clean-architecture/domain"
	"go-clean-architecture/domain/responses"
	"net/http"

	"github.com/gofiber/fiber"
)

type MetricController struct {
	MetricUsecase domain.MetricUsecase
	Env           *bootstrap.Env
}

func (mc *MetricController) Create(c *fiber.Ctx) {
	var metric domain.Metric

	err := c.BodyParser(&metric)
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(responses.ErrorResponse{Message: err.Error()})
		return
	}

	if err != nil {
		c.Status(http.StatusBadRequest).JSON(responses.ErrorResponse{Message: err.Error()})
		return
	}

	err = mc.MetricUsecase.Create(c.Context(), &metric)
	if err != nil {
		c.Status(http.StatusInternalServerError).JSON(responses.ErrorResponse{Message: err.Error()})
		return
	}

	c.Status(http.StatusOK).JSON(responses.SuccessResponse{
		Message: "Task created successfully",
	})
}

func (mc *MetricController) Fetch(c *fiber.Ctx) {
	metrics, err := mc.MetricUsecase.Fetch(c.Context())
	if err != nil {
		c.Status(http.StatusInternalServerError).JSON(responses.ErrorResponse{Message: err.Error()})
		return
	}

	c.Status(http.StatusOK).JSON(metrics)
}

func (mc *MetricController) GetByID(c *fiber.Ctx) {
	id := c.Params(":id")

	metric, err := mc.MetricUsecase.GetByID(c.Context(), id)
	if err != nil {
		c.Status(http.StatusInternalServerError).JSON(responses.ErrorResponse{Message: err.Error()})
		return
	}

	c.Status(http.StatusOK).JSON(metric)
}
