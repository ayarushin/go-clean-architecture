package domain

import "context"

const (
	MetricTable = "metrics"
)

type Metric struct {
	ID   uint64
	Name string
}

type MetricRepository interface {
	Create(c context.Context, metric *Metric) error
	Fetch(c context.Context, conds ...interface{}) ([]Metric, error)
	GetByID(c context.Context, id string) (Metric, error)
}

type MetricUsecase interface {
	Create(c context.Context, metric *Metric) error
	Fetch(c context.Context, conds ...interface{}) ([]Metric, error)
	GetByID(c context.Context, id string) (Metric, error)
}
