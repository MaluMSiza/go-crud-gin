package main

import (
	"context"
	"github.com/MaLuMSiza/go-crud-gin/internal/config"
	"github.com/MaLuMSiza/go-crud-gin/internal/handler"
	"github.com/MaLuMSiza/go-crud-gin/internal/repository"
	"github.com/MaLuMSiza/go-crud-gin/internal/routes"
	"github.com/MaLuMSiza/go-crud-gin/internal/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"log"
)

func main() {
	app := fx.New(
		fx.Provide(
			config.NewDatabase,
			newGinEngine,
			repository.NewUserRepository,
			service.NewUserService,
			handler.NewUserHandler,
		),
		fx.Invoke(registerRoutes),
	)

	if err := app.Start(context.Background()); err != nil {
		log.Fatal(err)
	}
}
func newGinEngine() *gin.Engine {
	return gin.Default()
}
func registerRoutes(router *gin.Engine, userHandler *handler.UserHandler) {
	routes.RegisterRoutes(router, userHandler)
	router.Run(":8080")
}
