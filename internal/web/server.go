package web

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"program_service/config"
	"program_service/internal/controllers"
	"program_service/internal/web/middlewares"
)

type WebServer struct {
	cfg    *config.WebConfig
	client *echo.Echo
}

func New(cfg *config.WebConfig) *WebServer {
	return &WebServer{
		cfg:    cfg,
		client: echo.New(),
	}
}

func (s *WebServer) RegisterRoutes(routes []controllers.Controller) {
	for _, route := range routes {
		mainGroup := s.client.Group("/api")
		group := mainGroup.Group(route.GetGroup())
		for _, handler := range route.GetHandlers() {
			switch handler.GetMethod() {
			case "GET":
				group.GET(handler.GetPath(), handler.GetHandler())
			case "POST":
				group.POST(handler.GetPath(), handler.GetHandler())
			case "PUT":
				group.PUT(handler.GetPath(), handler.GetHandler())
			case "PATCH":
				group.PATCH(handler.GetPath(), handler.GetHandler())
			case "DELETE":
				group.DELETE(handler.GetPath(), handler.GetHandler())
			default:
				panic("unsupported method")
			}
		}
	}

}

func (s *WebServer) RegisterMiddlewares(middlewares []middlewares.Middleware) {
	for _, middleware := range middlewares {
		switch middleware.GetLocation() {
		case "before":
			s.client.Pre(middleware.GetHandler())
		case "after":
			s.client.Use(middleware.GetHandler())
		default:
			panic("unsupported middleware location")
		}
	}
}

func (s *WebServer) Listen() {
	s.client.Start(fmt.Sprintf("%s:%s", s.cfg.Host, s.cfg.Port))
}

func (s *WebServer) Release(ctx context.Context) error {
	return s.client.Shutdown(ctx)
}
