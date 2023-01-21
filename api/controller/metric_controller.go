package controller

import (
	"go-clean-architecture/bootstrap"
	"go-clean-architecture/domain"
	"go-clean-architecture/domain/responses"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MetricController struct {
	MetricUsecase domain.MetricUsecase
	Env           *bootstrap.Env
}

func (mc *MetricController) Create(c *gin.Context) {
	var metric domain.Metric

	err := c.ShouldBind(&metric)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Message: err.Error()})
		return
	}

	id, err := strconv.ParseUint(c.GetString("x-metric-id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Message: err.Error()})
		return
	}

	metric.ID = id

	err = mc.MetricUsecase.Create(c, &metric)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, responses.SuccessResponse{
		Message: "Task created successfully",
	})
}

func (u *MetricController) Fetch(c *gin.Context) {
	metrics, err := u.MetricUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, metrics)
}

func (u *MetricController) GetByID(c *gin.Context) {
	id := c.GetString("x-metric-id")

	metric, err := u.MetricUsecase.GetByID(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, metric)
}
