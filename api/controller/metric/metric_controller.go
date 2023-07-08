package metric

import (
	"go-clean-architecture/bootstrap/env"
	"go-clean-architecture/domain"
	"go-clean-architecture/domain/responses"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type controller struct {
	Usecase domain.MetricUsecase
	Env     *env.Env
}

func New(usecase domain.MetricUsecase, env *env.Env) *controller {
	return &controller{
		Usecase: usecase,
		Env:     env,
	}
}

func (c *controller) Create(ctx *fiber.Ctx) error {
	var metric domain.Metric

	err := ctx.BodyParser(&metric)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(responses.ErrorResponse{Message: err.Error()})
	}

	err = c.Usecase.Create(ctx.Context(), &metric)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(responses.ErrorResponse{Message: err.Error()})
	}

	return ctx.Status(http.StatusOK).JSON(responses.SuccessResponse{
		Message: "Metric created successfully",
	})
}

func (c *controller) Fetch(ctx *fiber.Ctx) error {
	metrics, err := c.Usecase.Fetch(ctx.Context())
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(responses.ErrorResponse{Message: err.Error()})
	}

	return ctx.Status(http.StatusOK).JSON(metrics)
}

func (c *controller) GetByID(ctx *fiber.Ctx) error {
	id := ctx.Params(":id")

	metric, err := c.Usecase.GetByID(ctx.Context(), id)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(responses.ErrorResponse{Message: err.Error()})
	}

	return ctx.Status(http.StatusOK).JSON(metric)
}
