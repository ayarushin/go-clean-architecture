package controller

type Context interface {
	JSON(code int, obj any)
	Bind(i interface{}) error
}
