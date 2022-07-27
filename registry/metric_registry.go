package registry

func (r *registry) NewMetricController() MetricController {
	return NewMetricController(r.newMetricInteractor())
}

func (r *registry) newMetricInteractor() MetricInteractor {
	return NewMetricInteractor(r.newMetricRepository(), r.newMetricPresenter())
}

func (r *registry) newMetricRepository() MetricRepository {
	return NewMetricRepository(r.db)
}

func (r *registry) newMetricPresenter() MetricPresenter {
	return NewMetricPresenter()
}
