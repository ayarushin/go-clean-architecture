package bootstrap

import (
	"go-clean-architecture/bootstrap/internal/database"
	"go-clean-architecture/bootstrap/internal/env"
	"time"

	"gorm.io/gorm"
)

type ApplicationBuilder struct {
	env     *env.Env
	db      *gorm.DB
	timeout time.Duration
}

type Application struct {
	Env     *env.Env
	Db      *gorm.DB
	Timeout time.Duration
}

func Builder() *ApplicationBuilder {
	return &ApplicationBuilder{}
}

func (b *ApplicationBuilder) WithEnv() *ApplicationBuilder {
	b.env = env.New()
	return b
}

func (b *ApplicationBuilder) WithDB() *ApplicationBuilder {
	b.db = database.New(b.env)
	return b
}

func (b *ApplicationBuilder) WithTimeout() *ApplicationBuilder {
	b.timeout = time.Duration(b.env.ContextTimeout) * time.Second
	return b
}

func (b *ApplicationBuilder) Build() *Application {
	app := &Application{
		Env:     b.env,
		Db:      b.db,
		Timeout: b.timeout,
	}
	return app
}
