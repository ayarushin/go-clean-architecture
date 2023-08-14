package metric

import (
	"context"
	"time"

	"github.com/ayarushin/go-clean-architecture/domain"
)

type usecase struct {
	repository domain.MetricRepository
	timeout    time.Duration
}

func New(repository domain.MetricRepository, timeout time.Duration) domain.MetricUsecase {
	return &usecase{
		repository: repository,
		timeout:    timeout,
	}
}

func (u *usecase) Create(c context.Context, task *domain.Metric) error {
	ctx, cancel := context.WithTimeout(c, u.timeout)
	defer cancel()
	return u.repository.Create(ctx, task)
}

func (u *usecase) Update(c context.Context, task *domain.Metric) error {
	ctx, cancel := context.WithTimeout(c, u.timeout)
	defer cancel()
	return u.repository.Update(ctx, task)
}

func (u *usecase) Fetch(c context.Context, conds ...interface{}) ([]domain.Metric, error) {
	ctx, cancel := context.WithTimeout(c, u.timeout)
	defer cancel()
	return u.repository.Fetch(ctx, conds)
}

func (u *usecase) GetByID(c context.Context, id string) (domain.Metric, error) {
	ctx, cancel := context.WithTimeout(c, u.timeout)
	defer cancel()
	return u.repository.GetByID(ctx, id)
}
