package bootstrap

import (
	"go-clean-architecture/bootstrap/database"
	"go-clean-architecture/bootstrap/env"

	"gorm.io/gorm"
)

type Application struct {
	Env *env.Env
	Db  *gorm.DB
}

func New() Application {
	app := &Application{}
	app.Env = env.New()
	app.Db = database.New(app.Env)
	return *app
}
