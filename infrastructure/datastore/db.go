package datastore

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func NewDb() *gorm.DB {
	dsn := "host=localhost user=postgres password=default dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	return db
}
