package database

import (
	"fmt"
	"log"

	"github.com/ayarushin/go-clean-architecture/bootstrap/internal/env"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New(env *env.Env) *gorm.DB {
	dbHost := env.DBHost
	dbPort := env.DBPort
	dbUser := env.DBUser
	dbPass := env.DBPass
	dbName := env.DBName

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPass, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	return db
}
