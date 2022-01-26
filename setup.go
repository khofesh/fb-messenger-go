package main

import (
	"fb-messenger-go/controllers"

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
	router.GET("/ping", controllers.Ping)

	router.GET("/webhook", controllers.CheckToken)
	router.POST("/webhook", controllers.Callback)

	// 404
	router.NoRoute(controllers.HandleNotFound)

	return router
}
