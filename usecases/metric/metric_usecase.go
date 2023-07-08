package metric

import (
	"context"
	"go-clean-architecture/domain"
	"time"
)

type usecase struct {
	metricRepository domain.MetricRepository
	contextTimeout   time.Duration
}

func New(metricRepository domain.MetricRepository, timeout time.Duration) domain.MetricUsecase {
	return &usecase{
		metricRepository: metricRepository,
		contextTimeout:   timeout,
	}
}

func (u *usecase) Create(c context.Context, task *domain.Metric) error {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	return u.metricRepository.Create(ctx, task)
}

func (u *usecase) Fetch(c context.Context, conds ...interface{}) ([]domain.Metric, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	return u.metricRepository.Fetch(ctx, conds)
}

func (u *usecase) GetByID(c context.Context, id string) (domain.Metric, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	return u.metricRepository.GetByID(ctx, id)
}
