package controllers

import "github.com/labstack/echo/v4"

type Controller interface {
	GetGroup() string
	GetHandlers() []ControllerHandler
}

type ControllerHandler interface {
	GetMethod() string
	GetPath() string
	GetHandler() func(c echo.Context) error
}

type Handler struct {
	Method  string
	Path    string
	Handler func(c echo.Context) error
}

func (h *Handler) GetPath() string {
	return h.Path
}

func (h *Handler) GetHandler() func(c echo.Context) error {
	return h.Handler
}

func (h *Handler) GetMethod() string {
	return h.Method
}
