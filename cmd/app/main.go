package main

import (
	"context"
	"github.com/jmoiron/sqlx"
	"os"
	"os/signal"
	"program_service/config"
	"program_service/internal/controllers"
	"program_service/internal/controllers/program"
	"program_service/internal/database"
	"program_service/internal/web"
	"program_service/internal/web/middlewares"
	"program_service/libs"
	"sync"
)

func main() {
	logger := libs.NewApiLogger()
	logger.InitLogger()
	cfg, err := config.InitConfig()
	if err != nil {
		logger.Fatal(err.Error())
	}

	dbClient, err := database.New(cfg.DB)
	logger.Info("Connecting to database...")
	if err != nil {
		logger.Fatal(err.Error())
	}

	server := web.New(cfg.Web)
	RegisterRoutes(server, dbClient, logger)
	RegisterMiddlewares(server)

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		logger.Info("Starting server...")
		server.Listen()
	}()

	// GRACEFUL shutdown
	go func() {
		defer wg.Done()
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt)
		<-quit
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		if err = server.Release(ctx); err != nil {
			logger.Error(err.Error())
		}
	}()

	wg.Wait()
}

func RegisterRoutes(webClient *web.WebServer, dbClient *sqlx.DB, logger libs.Logger) {

	routes := []controllers.Controller{
		program.New(dbClient, logger, "https://example.com/api/programs/"),
	}
	webClient.RegisterRoutes(routes)
}
func RegisterMiddlewares(webClient *web.WebServer) {
	mws := []middlewares.Middleware{
		//CORS
		//AccessControl
	}
	webClient.RegisterMiddlewares(mws)
}
