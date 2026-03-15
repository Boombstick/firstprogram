package router

import (
	"firstprogram/handlers"
	"firstprogram/services"

	_ "firstprogram/docs"

	"github.com/gin-gonic/gin"
	ginFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Router struct {
	userService   services.IUserService
	counteService services.ICounterService
}

func New(userService services.IUserService, counterService services.ICounterService) *Router {
	return &Router{
		userService:   userService,
		counteService: counterService}
}

func (router *Router) SetupRoutes() *gin.Engine {

	r := gin.Default()

	r.POST("/sign/hmacsha512", handlers.SignHandler())
	r.POST("/users/create", handlers.CreateUserHandler(router.userService))
	r.POST("/counter/incr", handlers.CounterIncrHandler(router.counteService))

	r.GET("/swagger/*any", ginSwagger.WrapHandler(ginFiles.Handler))

	return r
}
