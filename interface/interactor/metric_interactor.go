package interactor

import "go-clean-architecture/domain/model"

type MetricInteractor interface {
	Get(m []*model.Metric) ([]*model.Metric, error)
}
