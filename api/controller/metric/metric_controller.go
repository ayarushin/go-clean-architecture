package metric

import (
	"net/http"

	"github.com/ayarushin/go-clean-architecture/domain"
	"github.com/ayarushin/go-clean-architecture/domain/responses"

	"github.com/gofiber/fiber/v2"
)

type controller struct {
	usecase domain.MetricUsecase
}

func New(usecase domain.MetricUsecase) *controller {
	return &controller{
		usecase: usecase,
	}
}

func (c *controller) Create(ctx *fiber.Ctx) error {
	var metric domain.Metric

	err := ctx.BodyParser(&metric)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(responses.ErrorResponse{Message: err.Error()})
	}

	err = c.usecase.Create(ctx.Context(), &metric)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(responses.ErrorResponse{Message: err.Error()})
	}

	return ctx.Status(http.StatusOK).JSON(responses.SuccessResponse{
		Message: "Metric created successfully",
	})
}

func (c *controller) Update(ctx *fiber.Ctx) error {
	var metric domain.Metric

	err := ctx.BodyParser(&metric)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(responses.ErrorResponse{Message: err.Error()})
	}

	err = c.usecase.Update(ctx.Context(), &metric)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(responses.ErrorResponse{Message: err.Error()})
	}

	return ctx.Status(http.StatusOK).JSON(responses.SuccessResponse{
		Message: "Metric updated successfully",
	})
}

func (c *controller) Fetch(ctx *fiber.Ctx) error {
	metrics, err := c.usecase.Fetch(ctx.Context())
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(responses.ErrorResponse{Message: err.Error()})
	}

	return ctx.Status(http.StatusOK).JSON(metrics)
}

func (c *controller) GetByID(ctx *fiber.Ctx) error {
	id := ctx.Params(":id")

	metric, err := c.usecase.GetByID(ctx.Context(), id)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(responses.ErrorResponse{Message: err.Error()})
	}

	return ctx.Status(http.StatusOK).JSON(metric)
}
