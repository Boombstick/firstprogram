package main

import (
	"context"
	"firstprogram/app"
	"firstprogram/cache"
	"firstprogram/config"
	"firstprogram/database"
	"firstprogram/repositories"
	"firstprogram/services"
	"fmt"
	"log"
	"net/http"
)

// @title        MyService API
// @version      1.0
// @description  HTTP-сервис с Redis, HMAC и PostgreSQL
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
	application := app.New(pgService, redisService)

	router := application.SetupRoutes()
	fmt.Printf("Запуск http сервера на порту %s", cfg.ServerPort)

	err = http.ListenAndServe(":"+cfg.ServerPort, router)
	if err != nil {
		log.Fatal("Не удалось запустить сервер: ", err)
	}
}
