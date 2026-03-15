package main

import (
	"context"
	"firstprogram/cache"
	"firstprogram/config"
	"firstprogram/database"
	"firstprogram/repositories"
	"firstprogram/router"
	"firstprogram/services"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

// @title        MyService API
// @version      1.0
// @host         localhost:8080
// @BasePath     /
func main() {

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки файлов конфигурации: ", err)
	}

	db, err := database.NewPostgres(database.PgConfig{
		Host:     cfg.PostgresHost,
		Port:     cfg.PostgresPort,
		User:     cfg.PostgresUser,
		Password: cfg.PostgresPassword,
		Database: cfg.PostgresDB,
	}, logger)
	if err != nil {
		logger.Fatal("ошибка подключения к PostgreSQL", zap.Error(err))
	}

	userRepo, err := repositories.NewUserRepository(context.Background(), db, logger)
	if err != nil {
		logger.Fatal("ошибка инициализации репозитория", zap.Error(err))
	}

	redisCache, err := cache.NewRedisCache(cfg.RedisHost, cfg.RedisPort, logger)
	if err != nil {
		logger.Fatal("ошибка подключения к Redis", zap.Error(err))
	}

	pgService := services.NewPostgresService(userRepo, logger)
	redisService := services.NewRedisService(redisCache, logger)
	router := router.New(pgService, redisService)

	httpHandler := router.SetupRoutes()

	server := &http.Server{
		Addr:    ":" + cfg.ServerPort,
		Handler: httpHandler,
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	g, gCtx := errgroup.WithContext(ctx)

	g.Go(func() error {
		logger.Info("сервер запущен", zap.String("port", cfg.ServerPort))
		return server.ListenAndServe()
	})

	g.Go(func() error {
		<-gCtx.Done()
		logger.Info("завершение сервера...")
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		return server.Shutdown(shutdownCtx)
	})

	if err := g.Wait(); err != nil {
		logger.Error("сервер остановлен с ошибкой", zap.Error(err))
	}
}
