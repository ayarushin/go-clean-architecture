package usecases

import (
	"context"
	"go-clean-architecture/domain"
	"time"
)

type metricUsecase struct {
	metricRepository domain.MetricRepository
	contextTimeout   time.Duration
}

func NewMetricUsecase(metricRepository domain.MetricRepository, timeout time.Duration) domain.MetricUsecase {
	return &metricUsecase{
		metricRepository: metricRepository,
		contextTimeout:   timeout,
	}
}

func (mu *metricUsecase) Create(c context.Context, task *domain.Metric) error {
	ctx, cancel := context.WithTimeout(c, mu.contextTimeout)
	defer cancel()
	return mu.metricRepository.Create(ctx, task)
}

func (mu *metricUsecase) Fetch(c context.Context, conds ...interface{}) ([]domain.Metric, error) {
	ctx, cancel := context.WithTimeout(c, mu.contextTimeout)
	defer cancel()
	return mu.metricRepository.Fetch(ctx, conds)
}

func (mu *metricUsecase) GetByID(c context.Context, id string) (domain.Metric, error) {
	ctx, cancel := context.WithTimeout(c, mu.contextTimeout)
	defer cancel()
	return mu.metricRepository.GetByID(ctx, id)
}
