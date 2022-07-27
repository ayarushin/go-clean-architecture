package interactor

import (
	"go-clean-architecture/domain/model"
	"go-clean-architecture/interface/interactor"
	"go-clean-architecture/interface/presenter"
	"go-clean-architecture/interface/repository"
)

type metricInteractor struct {
	metricRepository repository.MetricRepository
	metricPresenter  presenter.MetricPresenter
}

func NewMetricInteractor(r repository.MetricRepository, p presenter.MetricPresenter) interactor.MetricInteractor {
	return &metricInteractor{r, p}
}

func (mi *metricInteractor) Get(m []*model.Metric) ([]*model.Metric, error) {
	m, err := mi.metricRepository.FindAll(m)
	if err != nil {
		return nil, err
	}

	return mi.metricPresenter.ResponseMetrics(m), nil
}
