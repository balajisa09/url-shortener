package main

import (
	"encoding/json"
	"io/ioutil"
	"github.com/balajisa09/url-shortener/pkg/apis/shortener"
	"github.com/gin-gonic/gin"
)

var ServerURL string
type Request struct{
	URL string `json:"url"`
}

type Response struct{
	URL string `json:"url"`
	Hash string `json:"hash"`
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
	ctx.IndentedJSON(200,ServerURL+result)
}

func getURLHash(ctx *gin.Context){
	
	var respList []Response	
	urlhashmap := shortener.UrlHashMap
	for key, value := range urlhashmap{
		resp := Response{
			 URL: key,
			 Hash: value,
		}
		respList = append(respList, resp)
	}
	ctx.IndentedJSON(200,respList)
}

func main(){

	ServerURL = "https://short.in/"
	router := gin.Default()
	router.GET("/healthcheck",healthCheck)
	router.POST("/short",short)
	router.GET("/admin",getURLHash)
	router.Run(":8080")
}