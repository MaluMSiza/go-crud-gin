package routes

import (
	"github.com/MaLuMSiza/go-crud-gin/internal/handler"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, userHandler *handler.UserHandler) {
	router.GET("/users", userHandler.GetUsers)
	router.POST("/users", userHandler.CreateUser)
}
