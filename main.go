package main

import (
	"firstprogram/app"
	"firstprogram/cache"
	"firstprogram/config"
	"firstprogram/repositories"
	"firstprogram/services"
	"fmt"
	"log"
	"net/http"

	"github.com/go-pg/pg/v10"
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

	db := pg.Connect(&pg.Options{
		Addr:     cfg.PostgresHost + ":" + cfg.PostgresPort,
		User:     cfg.PostgresUser,
		Password: cfg.PostgresPassword,
		Database: cfg.PostgresDB})
	if _, err := db.Exec("SELECT 1"); err != nil {
		log.Fatal("не удалось подключиться к PostgreSQL: %w", err)
	}
	fmt.Println("PostgreSQL: подключено")

	userRepo := repositories.NewUserRepository(db)
	if err := userRepo.CreateTable(); err != nil {
		log.Fatal("Ошибка создания таблицы: ", err)
	}
	fmt.Println("PostgreSQL: таблица users готова")

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
