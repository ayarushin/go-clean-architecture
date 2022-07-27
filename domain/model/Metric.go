package model

type Metric struct {
	ID     uint   `gorm:"primary_key" json:"id"`
	Name   string `json:"name"`
	UnitId string `json:"unitId"`
}

func (Metric) TableName() string { return "metric" }
