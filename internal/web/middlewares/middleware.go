package middlewares

import "github.com/labstack/echo/v4"

type MiddlewareHandler struct {
	Location string
	Handler  func(next echo.HandlerFunc) echo.HandlerFunc
}

type Middleware interface {
	GetLocation() string
	GetHandler() func(next echo.HandlerFunc) echo.HandlerFunc
}
