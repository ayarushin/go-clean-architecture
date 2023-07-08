package metric

import (
	"context"
	"go-clean-architecture/domain"

	"gorm.io/gorm"
)

type repository struct {
	database *gorm.DB
	table    string
}

func New(db *gorm.DB, table string) domain.MetricRepository {
	return &repository{
		database: db,
		table:    table,
	}
}

func (mr *repository) Create(c context.Context, metric *domain.Metric) error {
	table := mr.database.WithContext(c).Table(mr.table)
	tx := table.Create(metric)
	return tx.Error
}

func (mr *repository) Fetch(c context.Context, conds ...interface{}) ([]domain.Metric, error) {
	table := mr.database.WithContext(c).Table(mr.table)
	var metrics []domain.Metric

	tx := table.Find(&metrics, conds)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return metrics, nil
}

func (mr *repository) GetByID(c context.Context, id string) (domain.Metric, error) {
	table := mr.database.WithContext(c).Table(mr.table)

	var metric domain.Metric

	tx := table.First(&metric, "id = ?", id)

	if tx.Error != nil {
		return domain.Metric{}, tx.Error
	}

	return metric, nil
}
