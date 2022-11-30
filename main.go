package main

import (
	"github.com/balajisa09/url-shortener/handlers"
	"github.com/gin-gonic/gin"
)


func main(){
	router := gin.Default()

	//register the apis to the router
	router.StaticFile("/","templates/index.html")
	router.GET("/healthcheck",handlers.HealthCheck)
	router.POST("/short",handlers.ShortenURL)
	router.GET("/admin",handlers.GetURLHashMap)
	router.GET("/shortner",handlers.RedirectURL)
	//start the server and listen on port 8080
	router.Run(":8080")
}