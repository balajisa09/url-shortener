package main

import (
	"github.com/balajisa09/url-shortener/handlers"
	"github.com/gin-gonic/gin"
)


func main(){
	router := gin.Default()
	router.GET("/healthcheck",handlers.HealthCheck)
	router.POST("/short",handlers.Short)
	router.GET("/admin",handlers.GetURLHash)
	router.Run(":8080")
}