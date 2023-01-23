package repository

import (
	"context"
	"go-clean-architecture/domain"

	"gorm.io/gorm"
)

type metricRepository struct {
	database *gorm.DB
	table    string
}

func NewMetricRepository(db *gorm.DB, table string) domain.MetricRepository {
	return &metricRepository{
		database: db,
		table:    table,
	}
}

func (mr *metricRepository) Create(c context.Context, metric *domain.Metric) error {
	table := mr.database.Table(mr.table)
	tx := table.Create(metric)
	return tx.Error
}

func (mr *metricRepository) Fetch(c context.Context, conds ...interface{}) ([]domain.Metric, error) {
	table := mr.database.Table(mr.table)
	var metrics []domain.Metric

	tx := table.Find(&metrics, conds)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return metrics, nil
}

func (mr *metricRepository) GetByID(c context.Context, id string) (domain.Metric, error) {
	table := mr.database.Table(mr.table)

	var metric domain.Metric

	tx := table.First(&metric, "id = ?", id)

	if tx.Error != nil {
		return domain.Metric{}, tx.Error
	}

	return metric, nil
}
