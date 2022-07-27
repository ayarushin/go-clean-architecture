package repository

import (
	"go-clean-architecture/domain/model"
	"go-clean-architecture/interface/repository"
	"gorm.io/gorm"
)

type metricRepository struct {
	db *gorm.DB
}

func NewMetricRepository(db *gorm.DB) repository.MetricRepository {
	return &metricRepository{db}
}

func (mr *metricRepository) FindAll(m []*model.Metric) ([]*model.Metric, error) {
	err := mr.db.Find(&m).Error

	if err != nil {
		return nil, err
	}

	return m, nil
}
