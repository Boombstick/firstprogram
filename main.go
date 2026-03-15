package main

import (
	"context"
	"firstprogram/cache"
	"firstprogram/config"
	"firstprogram/database"
	"firstprogram/repositories"
	"firstprogram/router"
	"firstprogram/services"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

// @title        MyService API
// @version      1.0
// @host         localhost:8080
// @BasePath     /
func main() {

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
	})
	if err != nil {
		log.Fatal(err)
	}

	userRepo, err := repositories.NewUserRepository(context.Background(), db)
	if err != nil {
		log.Fatal(err)
	}

	redisCache, err := cache.NewRedisCache(cfg.RedisHost, cfg.RedisPort)
	if err != nil {
		log.Fatal("ошибка подключения к Redis: ", err)
	}

	pgService := services.NewPostgresService(userRepo)
	redisService := services.NewRedisService(redisCache)
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
		fmt.Printf("Запуск http сервера на порту %s", cfg.ServerPort)
		return server.ListenAndServe()
	})

	g.Go(func() error {
		<-gCtx.Done()
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		return server.Shutdown(shutdownCtx)
	})

	if err := g.Wait(); err != nil {
		fmt.Printf("exit reason: %s \n", err)
	}
}
