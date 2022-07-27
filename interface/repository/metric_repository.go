package repository

import "go-clean-architecture/domain/model"

type MetricRepository interface {
	FindAll(m []*model.Metric) ([]*model.Metric, error)
}
