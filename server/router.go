package server

import (
	"github.com/gin-gonic/gin"
	"github.com/rodgomes/go-quickstart/controllers"
	"github.com/rodgomes/go-quickstart/middlewares"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	health := new(controllers.HealthController)

	router.GET("/health", health.Status)
	router.Use(middlewares.AuthMiddleware())
	return router

}
