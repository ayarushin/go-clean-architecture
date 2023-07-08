package app

import (
	"go-clean-architecture/bootstrap/database"
	"go-clean-architecture/bootstrap/env"

	"gorm.io/gorm"
)

type ApplicationBuilder struct {
	env *env.Env
	db  *gorm.DB
}

type application struct {
	Env *env.Env
	Db  *gorm.DB
}

func NewBuilder() *ApplicationBuilder {
	return &ApplicationBuilder{}
}

func (b *ApplicationBuilder) WithEnv(env *env.Env) *ApplicationBuilder {
	b.env = env
	return b
}

func (b *ApplicationBuilder) WithDB() *ApplicationBuilder {
	b.db = database.New(b.env)
	return b
}

func (b *ApplicationBuilder) Build() *application {
	app := &application{
		Env: b.env,
		Db:  b.db,
	}
	return app
}
