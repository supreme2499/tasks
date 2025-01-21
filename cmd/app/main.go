package main

import (
	"context"

	"Tasks/internal/app"
	"Tasks/internal/config"
	"Tasks/internal/http-server/handlers"
	k "Tasks/internal/kafka"
	"Tasks/internal/lib/logger"
	"Tasks/internal/lib/logger/sl"
	repo "Tasks/internal/repository/postgres"
	repoCahe "Tasks/internal/repository/redis"
	"Tasks/internal/service"
	"Tasks/internal/storage"
)

func main() {
	cfg := config.MustLoad()

	log := logger.SetupLogger(cfg.Env)
	log.Info("start application")

	storages, err := storage.NewStorage(context.Background(), cfg)
	if err != nil {
		log.Error("failed to connect to database", sl.Err(err))
	}
	defer func() {
		if err := storages.Close(context.TODO()); err != nil {
			log.Error("Failed to close storages", sl.Err(err))
		}
	}()

	repoStorage := repo.NewStorage(storages.Postgres, log)
	repoCache := repoCahe.NewCache(storages.Redis, log)
	broker, err := k.New(cfg.KafkaAddresses)
	if err != nil {
		log.Error("failed to connect to kafka", sl.Err(err))
	}
	defer broker.Close()
	serv := service.NewService(repoStorage, repoCache, broker)
	deps := &handlers.Dependencies{
		Service: serv,
		Log:     log,
	}

	h := handlers.NewHandler(deps)
	router := app.SetupRouter(h, log)
	// TODO: разобраться с авторизацией
	server := app.New(cfg, log, router)
	if err := server.Run(); err != nil {
		log.Error("server stopped with error", sl.Err(err))
	}
}
