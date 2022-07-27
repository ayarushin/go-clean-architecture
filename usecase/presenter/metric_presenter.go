package presenter

import (
	"go-clean-architecture/domain/model"
	"go-clean-architecture/interface/presenter"
)

type metricPresenter struct{}

func NewMetricPresenter() presenter.MetricPresenter {
	return &metricPresenter{}
}

func (mp *metricPresenter) ResponseMetrics(m []*model.Metric) []*model.Metric {
	// Do something and return
	return m
}
