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
	Create(ctx context.Context, metric *Metric) error
	Fetch(ctx context.Context, conds ...interface{}) ([]Metric, error)
	GetByID(ctx context.Context, id string) (Metric, error)
}

type MetricUsecase interface {
	Create(ctx context.Context, metric *Metric) error
	Fetch(ctx context.Context, conds ...interface{}) ([]Metric, error)
	GetByID(ctx context.Context, id string) (Metric, error)
}
