package app

import (
	"firstprogram/handlers"
	"firstprogram/services"

	"github.com/gin-gonic/gin"
	ginFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "firstprogram/docs"
)

type App struct {
	PostgresService *services.PostgresService
	RedisService    *services.RedisService
}

func New(pgService *services.PostgresService, redisService *services.RedisService) *App {
	return &App{
		PostgresService: pgService,
		RedisService:    redisService,
	}
}

func (app *App) SetupRoutes() *gin.Engine {

	r := gin.Default()

	r.POST("/sign/hmacsha512", handlers.SignHandler())
	r.POST("/postgres/users", handlers.PostgresUsersHandler(app.PostgresService))
	r.POST("/redis/incr", handlers.RedisIncrHandler(app.RedisService))

	r.GET("/swagger/*any", ginSwagger.WrapHandler(ginFiles.Handler))

	return r
}
