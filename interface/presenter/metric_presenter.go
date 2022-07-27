package presenter

import "go-clean-architecture/domain/model"

type MetricPresenter interface {
	ResponseMetrics(m []*model.Metric) []*model.Metric
}
