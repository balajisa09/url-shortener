package handlers

import(
	"encoding/json"
	"io/ioutil"
	"github.com/balajisa09/url-shortener/pkg/apis/shortener"
	"github.com/gin-gonic/gin"
	"github.com/balajisa09/url-shortener/models"
)

var ServerURL string

func HealthCheck(ctx *gin.Context){
	ctx.IndentedJSON(200,"App is Healthy")
}

func Short(ctx *gin.Context){
	data, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.IndentedJSON(400,"Invalid Request")
	}
	var req models.Request
	json.Unmarshal(data,&req)
	link := req.URL
	result := shortener.Short(link)
	result = result[:6]
	ctx.IndentedJSON(200,ServerURL+result)
}

func GetURLHash(ctx *gin.Context){
	
	var respList []models.Response	
	urlhashmap := shortener.UrlHashMap
	for key, value := range urlhashmap{
		resp := models.Response{
			 URL: key,
			 Hash: value,
		}
		respList = append(respList, resp)
	}
	ctx.IndentedJSON(200,respList)
}

func init(){
	ServerURL = "https://short.in/"
}