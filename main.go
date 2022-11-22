package main


import (
	"github.com/gin-gonic/gin"
)

func healthCheck(ctx *gin.Context){
	ctx.IndentedJSON(200,"App is Healthy")
}

func main(){

	router := gin.Default()

	router.GET("/healthcheck",healthCheck)
	router.Run(":8080")
}