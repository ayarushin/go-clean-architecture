package controller

import (
	"github.com/gin-gonic/gin"
	"go-clean-architecture/domain/model"
	"go-clean-architecture/interface/controller"
	"go-clean-architecture/interface/interactor"
	"net/http"
)

type metricController struct {
	metricInteractor interactor.MetricInteractor
}

func NewMetricController(mi interactor.MetricInteractor) controller.MetricController {
	return &metricController{metricInteractor: mi}
}

func (mc *metricController) GetMetrics(c controller.Context) {
	var m []*model.Metric

	m, err := mc.metricInteractor.Get(m)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, m)
}
