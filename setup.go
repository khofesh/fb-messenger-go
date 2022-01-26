package main

import (
	"fb-messenger-go/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
 * initiate routers
 */
func SetupRouter() *gin.Engine {
	// Run Gin in Release mode
	// gin.SetMode(gin.ReleaseMode)
	var trusted = []string{"192.168.1.2", "127.0.0.1"}

	router := gin.Default()
	router.SetTrustedProxies(trusted)

	// Ping test
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	router.GET("/webhook", controllers.CheckToken)
	router.POST("/webhook", controllers.Callback)

	// 404
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not Found"})
	})

	return router
}
