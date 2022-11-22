package main

import (
	"encoding/json"
	"io/ioutil"
	"github.com/balajisa09/url-shortener/pkg/apis/shortener"
	"github.com/gin-gonic/gin"
)

type Request struct{
	URL string `json:"url"`
}

func healthCheck(ctx *gin.Context){
	ctx.IndentedJSON(200,"App is Healthy")
}

func short(ctx *gin.Context){
	data, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.IndentedJSON(400,"Invalid Request")
	}
	var req Request
	json.Unmarshal(data,&req)
	link := req.URL
	result := shortener.Short(link)
	result = result[:6]
	ctx.IndentedJSON(200,result)
}

func main(){

	router := gin.Default()
	router.GET("/healthcheck",healthCheck)
	router.POST("/short",short)
	router.Run(":8080")
}